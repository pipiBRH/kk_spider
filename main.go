package main

import (
	"log"
	"os"

	database "github.com/pipiBRH/kk_database"

	"github.com/pipiBRH/kk_spider/config"
	"github.com/pipiBRH/kk_spider/dal"
	"github.com/pipiBRH/kk_spider/request"
)

func main() {
	err := config.NewConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	data := request.NewSourceDataset()
	err = data.GetDataset(config.Config.GovDatasetURL)
	if err != nil {
		log.Fatal(err)
	}

	bikeInfo, err := data.ConvertedToElasticsearchDataFormat()
	if err != nil {
		log.Fatal(err)
	}

	err = database.InitElasticsearchConnection(
		config.Config.Elasticsearch.Host,
		config.Config.Elasticsearch.Port,
	)
	if err != nil {
		log.Fatal(err)
	}

	esDAL := dal.NewElasticsearchDAL(database.EsClient)
	_, err = esDAL.CreateYouBikeInfoByBulk(config.Config.Elasticsearch.Index["youbike_history"], bikeInfo)
	if err != nil {
		log.Fatal(err)
	}

	_, err = esDAL.UpdateYouBikeInfoByBulk(config.Config.Elasticsearch.Index["youbike"], bikeInfo)
	if err != nil {
		log.Fatal(err)
	}
}
