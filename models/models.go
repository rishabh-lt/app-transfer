package models

type DataTransferRequest struct {
	SourceBucket      string `json:"source_bucket"`
	DestinationBucket string `json:"destination_bucket"`
	SourcePath        string `json:"source_path"`
	SourceRegion      string `json:"source_region"`
	DestinationRegion string `json:"destination_region"`
}
