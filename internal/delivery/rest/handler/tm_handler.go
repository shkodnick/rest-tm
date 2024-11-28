package handler

import (
	"github.com/go-game-dev/rest-tm/internal/service/tm"
)

type TmHandler struct {
	TmSrv *tm.TmService
}

func NewTmHandler(TmSrv *tm.TmService) TmHandler {
	return TmHandler{
		TmSrv: TmSrv,
	}
}
