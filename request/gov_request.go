package request

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/pipiBRH/kk_spider/dal"

	"github.com/pipiBRH/kk_spider/config"
)

type (
	SourceDataset struct {
		Success bool         `json:"success"`
		Result  SourceResult `json:"result"`
	}

	SourceResult struct {
		ResourceID string         `json:"resource_id"`
		Limit      int            `json:"limit"`
		Total      int            `json:"total"`
		Fields     []SourceField  `json:"fields"`
		Records    []SourceRecord `json:"records"`
	}

	SourceField struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}

	SourceRecord struct {
		Sno     string `json:"sno"`
		Sna     string `json:"sna"`
		Tot     string `json:"tot"`
		Sbi     string `json:"sbi"`
		Sarea   string `json:"sarea"`
		Mday    string `json:"mday"`
		Lat     string `json:"lat"`
		Lon     string `json:"lng"`
		Ar      string `json:"ar"`
		Sareaen string `json:"sareaen"`
		Snaen   string `json:"snaen"`
		Aren    string `json:"aren"`
		Bemp    string `json:"bemp"`
		Act     string `json:"act"`
	}
)

func NewSourceDataset() *SourceDataset {
	return new(SourceDataset)
}

func (ds *SourceDataset) GetDataset(config *config.Config) error {
	targetContent, err := Curl(config.GovDatasetURL)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(targetContent, ds); err != nil {
		return err
	}

	if !ds.Success {
		return errors.New("Retrieve gov data failure")
	}

	return nil
}

func (ds *SourceDataset) ConvertedToElasticsearchDataFormat() ([]dal.YouBikeInfo, error) {

	bikeInfo := make([]dal.YouBikeInfo, len(ds.Result.Records))

	for i, d := range ds.Result.Records {
		lat, err := strconv.ParseFloat(strings.Trim(d.Lat, " "), 64)
		if err != nil {
			return nil, err
		}

		lon, err := strconv.ParseFloat(strings.Trim(d.Lon, " "), 64)
		if err != nil {
			return nil, err
		}

		sno, err := strconv.ParseInt(strings.Trim(d.Sno, " "), 10, 0)
		if err != nil {
			return nil, err
		}

		tot, err := strconv.ParseInt(strings.Trim(d.Tot, " "), 10, 0)
		if err != nil {
			return nil, err
		}

		sbi, err := strconv.ParseInt(strings.Trim(d.Sbi, " "), 10, 0)
		if err != nil {
			return nil, err
		}

		bemp, err := strconv.ParseInt(strings.Trim(d.Bemp, " "), 10, 0)
		if err != nil {
			return nil, err
		}

		act, err := strconv.ParseInt(strings.Trim(d.Act, " "), 10, 0)
		if err != nil {
			return nil, err
		}

		t, err := time.Parse("20060102150405", strings.Trim(d.Mday, " "))
		if err != nil {
			return nil, err
		}

		bikeInfo[i] = dal.YouBikeInfo{
			Sno:      int(sno),
			Sna:      d.Sna,
			Tot:      int(tot),
			Sbi:      int(sbi),
			Sarea:    d.Sarea,
			Mday:     t.Format("2006-01-02 15:04:05"),
			Ar:       d.Ar,
			Sareaen:  d.Sareaen,
			Snaen:    d.Snaen,
			Aren:     d.Aren,
			Bemp:     int(bemp),
			Act:      int(act),
			Location: []float32{float32(lon), float32(lat)},
		}
	}

	return bikeInfo, nil
}
