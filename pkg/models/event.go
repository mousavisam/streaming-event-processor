package models

import "time"

type Event struct {
	ID            string                 `json:"id"`
	SourceId      string                 `json:"source_id"`
	Type          string                 `json:"type"`
	Data          map[string]interface{} `json:"data"`
	SchemaVersion string                 `json:"schema_version"`
	CreatedAt     time.Time              `json:"created_at"`
}
