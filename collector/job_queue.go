package collector

// https://moonraker.readthedocs.io/en/latest/web_api/#retrieve-the-job-queue-status

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type MoonrakerJobQueueResponse struct {
	Result struct {
		QueuedJobs []MoonrakerQueuedJob `json:"queued_jobs"`
		QueueState string               `json:"queue_state"`
	} `json:"result"`
}

type MoonrakerQueuedJob struct {
	TimeInQueue float64 `json:"time_in_queue"`
}

func (c Collector) fetchMoonrakerJobQueue(klipperHost string, apiKey string) (*MoonrakerJobQueueResponse, error) {
	var url = "http://" + klipperHost + "/server/job_queue/status"
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

	var response MoonrakerJobQueueResponse

	err = json.Unmarshal(data, &response)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &response, nil
}
