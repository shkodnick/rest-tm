package tm

type TmService struct {
	TmRepository
}

func New(repo TmRepository) (*TmService, error) {
	return &TmService{TmRepository: repo}, nil
}
