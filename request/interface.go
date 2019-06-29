package request

import (
	"github.com/pipiBRH/kk_spider/dal"
)

type Spider interface {
	GetDataset(string) error
	ConvertedToElasticsearchDataFormat() ([]dal.YouBikeInfo, error)
}
