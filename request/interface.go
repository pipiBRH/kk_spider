package request

import (
	"github.com/pipiBRH/kk_spider/config"
	"github.com/pipiBRH/kk_spider/dal"
)

type Spider interface {
	GetDataset(*config.Config) error
	ConvertedToElasticsearchDataFormat() ([]dal.YouBikeInfo, error)
}
