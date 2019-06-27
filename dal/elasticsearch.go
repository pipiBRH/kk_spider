package dal

import (
	"fmt"

	"github.com/olivere/elastic"

	"github.com/pipiBRH/kk_spider/database"
)

type ElasticsearchDAL struct {
	Es *database.ElasticsearchConnection
}

func NewElasticsearchDAL(client *database.ElasticsearchConnection) *ElasticsearchDAL {
	return &ElasticsearchDAL{
		Es: client,
	}
}

type YouBikeInfo struct {
	Sno      int
	Sna      string
	Tot      int
	Sbi      int
	Sarea    string
	Mday     string
	Ar       string
	Sareaen  string
	Snaen    string
	Aren     string
	Bemp     int
	Act      int
	Location []float32
}

// TODO: Need chunk data to pervent bulk overload
func (es *ElasticsearchDAL) CreateYouBikeInfoByBulk(i string, t string, data []YouBikeInfo) (*elastic.BulkResponse, error) {
	bulkRequest := es.Es.Client.Bulk()

	for _, d := range data {
		id := fmt.Sprintf("%d-%s", d.Sno, d.Mday)
		req := elastic.NewBulkIndexRequest().Index(i).Type(t).Id(id).Doc(d)
		bulkRequest = bulkRequest.Add(req)
	}

	bulkResponse, err := bulkRequest.Do(es.Es.Ctx)
	if err != nil {
		return nil, err
	}

	return bulkResponse, nil
}
