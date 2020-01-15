package es

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

var (
	r map[string]interface{}
)

// ElasticSearchDB :
type ElasticSearchDB struct {
	Client *elasticsearch.Client
}

// NewAdapter :
func (sv *ElasticSearchDB) NewAdapter() error {
	var err error
	sv.Client, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Printf("[ERROR]NewAdapter:%+v\n", err)
		return err
	}
	return nil
}

// InsertOrUpdate :
func (sv *ElasticSearchDB) InsertOrUpdate(index string, documentID string, jsonString string) (string, string, int) {
	req := esapi.IndexRequest{
		Index:      index,                         // "test",
		DocumentID: documentID,                    // strconv.Itoa(i + 1),
		Body:       strings.NewReader(jsonString), // strings.NewReader(`{"title" : "` + title + `"}`),
		Refresh:    "true",
	}
	// Perform the request with the client.
	res, err := req.Do(context.Background(), sv.Client)
	if err != nil {
		log.Fatalf("[ERROR] getting response: %s", err)
	}
	defer res.Body.Close()
	//do result
	if res.IsError() {
		log.Printf("[ERROR][%s] Error indexing document ID=%v", res.Status(), documentID)
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("[ERROR] parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
			return res.Status(), (r["result"]).(string), int(r["_version"].(float64))
		}
	}
	return "", "", 0
}

// Search :
func (sv *ElasticSearchDB) Search(index string, jsonString string) (int, []interface{}, error) {
	// Use the helper methods of the client.
	res, err := sv.Client.Search(
		sv.Client.Search.WithContext(context.Background()),
		sv.Client.Search.WithIndex(index),                        // test
		sv.Client.Search.WithBody(strings.NewReader(jsonString)), // `{"query" : { "match" : { "title" : "test" } }}`
		sv.Client.Search.WithTrackTotalHits(true),
		sv.Client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("[ERROR]: %s", err)
		return -1, nil, err
	}
	defer res.Body.Close()
	// err
	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("[ERROR] parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[ERROR][%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return -2, nil, err
	}
	// json err
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("[ERROR] parsing the response body: %s", err)
		return -3, nil, err
	}
	// Print the response status, number of results, and request duration.
	//log.Printf(
	//	"[%s] %d hits; took: %dms",
	//	res.Status(),
	//	int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
	//	int(r["took"].(float64)),
	//)
	// Print the ID and document source for each hit.
	//for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
	//	log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	//}
	total := int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64))
	list := r["hits"].(map[string]interface{})["hits"].([]interface{})
	return total, list, nil
}
