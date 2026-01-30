package parametertype

import "context"

type Service interface {
	ListParametersTypes(ctx context.Context) ([]AccountingParameterType, error)
	FindParameterTypeById(ctx context.Context, id string) (AccountingParameterType, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) ListParametersTypes(ctx context.Context) ([]AccountingParameterType, error) {
	return s.repo.ListParametersTypes(ctx)
}

func (s *service) FindParameterTypeById(ctx context.Context, id string) (AccountingParameterType, error) {
	return s.repo.FindParameterTypeById(ctx, id)
}
