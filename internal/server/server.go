package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/izzatbey/soc-kafka-preprocess/internal/config"
	"github.com/izzatbey/soc-kafka-preprocess/internal/kafka"
	"github.com/izzatbey/soc-kafka-preprocess/internal/preprocess"
)

func Start(cfg *config.Config) {
	producer := kafka.NewProducerWithBroker(cfg.KafkaBroker)
	defer producer.Close()

	http.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusBadRequest)
			return
		}

		if len(body) == 0 {
			http.Error(w, "Empty body", http.StatusBadRequest)
			return
		}

		var logs []map[string]interface{}
		if body[0] == '[' {
			if err := json.Unmarshal(body, &logs); err != nil {
				http.Error(w, "Invalid JSON array", http.StatusBadRequest)
				return
			}
		} else {
			var single map[string]interface{}
			if err := json.Unmarshal(body, &single); err != nil {
				http.Error(w, "Invalid JSON object", http.StatusBadRequest)
				return
			}
			logs = append(logs, single)
		}

		for _, logData := range logs {
			raw, _ := json.Marshal(logData)
			processed := preprocess.ApplyPreprocessRules(string(raw), cfg.LogTag)

			if processed == "" {
				continue
			}

			var dynamic map[string]interface{}
			if err := json.Unmarshal([]byte(processed), &dynamic); err != nil {
				log.Printf("skipping invalid processed JSON: %v", err)
				continue
			}

			if dynamic == nil {
				dynamic = make(map[string]interface{})
			}

			dynamic["@timestamp"] = time.Now().Format(time.RFC3339Nano)

			finalJSON, _ := json.Marshal(dynamic)
			producer.Publish("raw-logs", finalJSON)
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok"}`))
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	log.Printf("✅ HTTP server listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
