package parametertype

import "context"

type Repository interface {
	ListParametersTypes(ctx context.Context) ([]AccountingParameterType, error)
	FindParameterTypeById(ctx context.Context, id string) (AccountingParameterType, error)
}
