package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/olivere/elastic"

	"github.com/pipiBRH/kk_spider/config"
)

type ElasticsearchConnection struct {
	Ctx    context.Context
	Client *elastic.Client
}

func InitElasticsearchConnection(config *config.Config) (*ElasticsearchConnection, error) {
	ctx := context.Background()
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	host := fmt.Sprintf("http://%s:%d", config.Elasticsearch.Host, config.Elasticsearch.Port)

	client, err := elastic.NewClient(
		elastic.SetErrorLog(errorlog),
		elastic.SetURL(host),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}

	return &ElasticsearchConnection{
		Ctx:    ctx,
		Client: client,
	}, nil
}
