package es

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDBES_NewAdapter(t *testing.T) {
	sv := new(ElasticSearchDB)
	sv.NewAdapter()
	if sv.Client != nil {
		log.Println(sv.Client.Info())
	}
}

func TestDBES_InsertOrUpdate(t *testing.T) {
	sv := new(ElasticSearchDB)
	sv.NewAdapter()
	if sv.Client == nil {
		log.Println("client nil")
		return
	}
	status, result, version := sv.InsertOrUpdate("test", "11", `{"title" : "test-110-124-10"}`)
	log.Println("status=", status)
	log.Println("result=", result)
	log.Println("version=", version)
}

/*
2020/01/10 17:58:48 err= <nil>
2020/01/10 17:58:48 v=map[_id:1 _index:test _score:0.09141083 _source:map[title:test-110-123] _type:_doc]
2020/01/10 17:58:48 v=map[_id:2 _index:test _score:0.09141083 _source:map[title:test-110-123] _type:_doc]
2020/01/10 17:58:48 v=map[_id:10 _index:test _score:0.08115276 _source:map[title:test-110-124-10] _type:_doc]
2020/01/10 17:58:48 v=map[_id:11 _index:test _score:0.08115276 _source:map[title:test-110-124-10] _type:_doc]
--- PASS: TestDBES_find (0.01s)
*/
func TestDBES_find(t *testing.T) {
	sv := new(ElasticSearchDB)
	sv.NewAdapter()
	if sv.Client == nil {
		log.Println("client nil")
		return
	}
	total, list, err := sv.Search("test", `{"query" : { "match" : { "title" : "test" } }}`)
	log.Println("total=", total)
	log.Println("err=", err)
	if total > 0 {
		for _, v := range list {
			log.Printf("v=%+v\n", v)
		}
	}
}

func TestDBES_find_login(t *testing.T) {
	sv := new(ElasticSearchDB)
	sv.NewAdapter()
	if sv.Client == nil {
		log.Println("client nil")
		return
	}
	total, list, err := sv.Search("user100w", `{"query" : { "match" : { "name" : "1099998@qq.com" } }}`)
	log.Println("total=", total)
	log.Println("err=", err)
	if total > 0 {
		for _, v := range list {
			log.Printf("v=%+v\n", v)
		}
	}
}

//--------------------------------------------------------------

func TestDB_esInfo_1(t *testing.T) {
	//"github.com/elastic/go-elasticsearch"
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
}

func TestDB_esInfo_more(t *testing.T) {
	cfg := elasticsearch.Config{
		Addresses: []string{ // 配置多个集群
			"http://localhost:9200",
			"http://localhost:9200",
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	log.Println("err:", err)
	//"github.com/elastic/go-elasticsearch"
	//es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
}

// https://www.infoq.cn/article/HVzMNKuYYmCkRTK-oZdp
func TestDB_esData_Inster(t *testing.T) {
	//"github.com/elastic/go-elasticsearch"
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(es.Info())
	// do
	for i, title := range []string{"Test One", "Test Two"} {
		req := esapi.IndexRequest{
			Index:      "test",
			DocumentID: strconv.Itoa(i + 1),
			Body:       strings.NewReader(`{"title" : "` + title + `"}`),
			Refresh:    "true",
		}
		// Perform the request with the client.
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		defer res.Body.Close()
		//do result
		if res.IsError() {
			log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				// Print the response status and indexed document version.
				log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			}
		}
	}

}

//var (
//	r  map[string]interface{}
//	wg sync.WaitGroup
//)

/*
GET test/_search
{"query" : { "match" : { "title" : "test" } }}
::log
=== RUN   TestDB_esData_find
2020/01/10 11:31:41 [200 OK] 2 hits; took: 5ms
2020/01/10 11:31:41  * ID=1, map[title:Test One]
2020/01/10 11:31:41  * ID=2, map[title:Test Two]
--- PASS: TestDB_esData_find (0.02s)
*/
func TestDB_esData_findList(t *testing.T) {
	//es
	es, _ := elasticsearch.NewDefaultClient()
	// search
	// Use the helper methods of the client.
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(strings.NewReader(`{"query" : { "match" : { "title" : "test" } }}`)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	defer res.Body.Close()
	//err
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}
	// json err
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	//info
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d hits; took: %dms",
		res.Status(),
		int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(r["took"].(float64)),
	)
	// info
	// Print the ID and document source for each hit.
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
}

/*
{
  "took" : 1,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 2,
      "relation" : "eq"
    },
    "max_score" : 0.18232156,
    "hits" : [
      {
        "_index" : "test",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.18232156,
        "_source" : {
          "title" : "Test One"
        }
      },
      {
        "_index" : "test",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 0.18232156,
        "_source" : {
          "title" : "Test Two"
        }
      }
    ]
  }
*/
