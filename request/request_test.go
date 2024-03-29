package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurlWithTestServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		content := []byte(`ok`)
		res.Write(content)
	}))

	defer server.Close()

	body, err := Curl(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, string(body), "ok")
}

func TestGovRequestGetDataSet(t *testing.T) {
	content := []byte(`{"success":true,"result":{"resource_id":"382000000A-000352-001","limit":2000,"total":556,"fields":[{"type":"text","id":"sno"}],"records":[{"sno":"1001","sna":"大鵬華城","tot":"38","sbi":"17","sarea":"新店區","mday":"20190630201817","lat":"24.99116","lng":"121.53398","ar":"新北市新店區中正路700巷3號","sareaen":"Xindian Dist.","snaen":"Dapeng Community","aren":"No. 3, Lane 700 Chung Cheng Road, Xindian District","bemp":"21","act":"1"}]}}`)

	tDs := &SourceDataset{
		Success: true,
		Result: SourceResult{
			ResourceID: "382000000A-000352-001",
			Limit:      2000,
			Total:      556,
			Fields: []SourceField{
				SourceField{
					Type: "text",
					ID:   "sno",
				},
			},
			Records: []SourceRecord{
				SourceRecord{

					Sno:     "1001",
					Sna:     "大鵬華城",
					Tot:     "38",
					Sbi:     "17",
					Sarea:   "新店區",
					Mday:    "20190630201817",
					Lat:     "24.99116",
					Lon:     "121.53398",
					Ar:      "新北市新店區中正路700巷3號",
					Sareaen: "Xindian Dist.",
					Snaen:   "Dapeng Community",
					Aren:    "No. 3, Lane 700 Chung Cheng Road, Xindian District",
					Bemp:    "21",
					Act:     "1",
				},
			},
		},
	}

	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write(content)
	}))
	defer server.Close()

	ds := NewSourceDataset()
	err := ds.GetDataset(server.URL)
	assert.NoError(t, err)
	assert.Equal(t, ds, tDs)
}

func TestGovRequestConvertedToElasticsearchDataFormat(t *testing.T) {
	ds := &SourceDataset{
		Success: true,
		Result: SourceResult{
			ResourceID: "382000000A-000352-001",
			Limit:      2000,
			Total:      556,
			Fields: []SourceField{
				SourceField{
					Type: "text",
					ID:   "sno",
				},
			},
			Records: []SourceRecord{
				SourceRecord{

					Sno:     "1001 ",
					Sna:     "大鵬華城",
					Tot:     "38",
					Sbi:     "17",
					Sarea:   "新店區",
					Mday:    "20190630201817",
					Lat:     "24.99116",
					Lon:     "121.53398",
					Ar:      "新北市新店區中正路700巷3號",
					Sareaen: "Xindian Dist.",
					Snaen:   "Dapeng Community",
					Aren:    "No. 3, Lane 700 Chung Cheng Road, Xindian District",
					Bemp:    "21",
					Act:     "1",
				},
			},
		},
	}

	data, err := ds.ConvertedToElasticsearchDataFormat()
	assert.NoError(t, err)
	assert.Equal(t, len(data), 1)
	assert.Equal(t, data[0].Sno, 1001)
	assert.Equal(t, data[0].Mday, "2019-06-30 20:18:17")
}
