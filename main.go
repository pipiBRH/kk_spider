package main

import (
	"log"
	"os"

	"github.com/pipiBRH/kk_spider/config"
	"github.com/pipiBRH/kk_spider/dal"
	"github.com/pipiBRH/kk_spider/database"
	"github.com/pipiBRH/kk_spider/request"
)

func main() {
	config, err := config.NewConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	data := request.NewSourceDataset()
	err = data.GetDataset(config)
	if err != nil {
		log.Fatal(err)
	}

	bikeInfo, err := data.ConvertedToElasticsearchDataFormat()
	if err != nil {
		log.Fatal(err)
	}

	esClient, err := database.InitElasticsearchConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	esDAL := dal.NewElasticsearchDAL(esClient)
	_, err = esDAL.CreateYouBikeInfoByBulk(config.Elasticsearch.Index, config.Elasticsearch.Type, bikeInfo)
	if err != nil {
		log.Fatal(err)
	}
}
