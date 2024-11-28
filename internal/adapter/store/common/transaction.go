package common

import (
	"context"
	"github.com/jmoiron/sqlx"

	"database/sql"
	"github.com/pkg/errors"
)

func TransactionWrapper(
	ctx context.Context,
	conn *sqlx.DB,
	wrappedFunc func(ctx context.Context, tx *sql.Tx) error,
) error {
	tx, err := conn.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "conn.BeginTx")
	}

	err = wrapForRecover(ctx, tx, wrappedFunc)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return errors.Wrapf(err, "rollback error: %s", rollbackErr.Error())
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return errors.Wrap(err, "tx.Commit")
	}

	return nil
}

func wrapForRecover(
	ctx context.Context,
	tx *sql.Tx,
	wrappedFunc func(ctx context.Context, tx *sql.Tx) error,
) (err error) {
	defer func() {
		if errPanic := recover(); errPanic != nil {
			err = errors.Errorf("panic: %s", errPanic)
		}
	}()

	return wrappedFunc(ctx, tx)
}
