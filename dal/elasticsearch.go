package dal

import (
	"github.com/pipiBRH/kk_database"
	"fmt"
	"strconv"
	"time"

	"github.com/olivere/elastic"

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
func (es *ElasticsearchDAL) CreateYouBikeInfoByBulk(index string, data []YouBikeInfo) (*elastic.BulkResponse, error) {
	bulkRequest := es.Es.Client.Bulk()

	for _, d := range data {
		t, err := time.Parse("2006-01-02 15:04:05", d.Mday)
		if err != nil {
			return nil, err
		}
		id := fmt.Sprintf("%d-%s", d.Sno, t.Format("20060102150405"))
		req := elastic.NewBulkIndexRequest().Index(index).Type("_doc").Id(id).Doc(d)
		bulkRequest = bulkRequest.Add(req)
	}

	bulkResponse, err := bulkRequest.Do(es.Es.Ctx)
	if err != nil {
		return nil, err
	}

	return bulkResponse, nil
}

// TODO: Need chunk data to pervent bulk overload
func (es *ElasticsearchDAL) UpdateYouBikeInfoByBulk(index string, data []YouBikeInfo) (*elastic.BulkResponse, error) {
	bulkRequest := es.Es.Client.Bulk()

	for _, d := range data {
		req := elastic.NewBulkIndexRequest().Index(index).Type("_doc").Id(strconv.Itoa(d.Sno)).Doc(d)
		bulkRequest = bulkRequest.Add(req)
	}

	bulkResponse, err := bulkRequest.Do(es.Es.Ctx)
	if err != nil {
		return nil, err
	}

	return bulkResponse, nil
}
