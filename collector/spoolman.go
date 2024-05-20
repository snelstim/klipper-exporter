package collector

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type SpoolmanHealthQueryResponse struct {
	Result struct {
		Status string `json:"status"`
	} `json:"result"`
}

func (c Collector) fetchSpoolmanHealth(spoolmanHost string, apiKey string) (*SpoolmanHealthQueryResponse, error) {
	var url = "http://" + spoolmanHost + "/health"
	log.Debug("Collecting metrics from " + url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if apiKey != "" {
		req.Header.Set("X-API-KEY", apiKey)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var response SpoolmanHealthQueryResponse

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &response, nil
}

type SpoolmanFindSpoolQueryResponse struct {
	Result struct {
		Id       bool `json:"id"`
		Filament []struct {
			Id          int `json:"id"`
			SpoolWeight int `json:"spool_weight"`
		} `json:"filament"`
	} `json:"result"`
}

func (c Collector) fetchSpoolmanFindSpool(spoolmanHost string, apiKey string) (*SpoolmanFindSpoolQueryResponse, error) {
	var url = "http://" + spoolmanHost + "/spool"
	log.Debug("Collecting metrics from " + url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if apiKey != "" {
		req.Header.Set("X-API-KEY", apiKey)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var response SpoolmanFindSpoolQueryResponse

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &response, nil
}
