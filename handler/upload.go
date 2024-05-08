package handler

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/borowiak-m/underwriting-engine/data"
	"github.com/google/uuid"
)

type CreateFileUploadRequest struct {
	FileType data.FileType  `json:"fileType"`
	Mapping  map[string]int `json:"mapping"`
}

type CreateFileUploadResponse struct {
	ID uuid.UUID `json:"id"`
}

func Upload(w http.ResponseWriter, r *http.Request) error {
	// TO DO QC checks
	subs, err := ingestSubsFromFile(r.Body)
	if err != nil {
		return err
	}
	ProcessSubscriptions(subs)
	return nil
}

func CreateFileUpload(w http.ResponseWriter, r *http.Request) error {
	var req CreateFileUploadRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil
	}
	fileUpload := data.FileUpload{
		ID:      uuid.New(),
		Type:    req.FileType,
		Mapping: req.Mapping,
	}
	fmt.Println(fileUpload)
	resp := CreateFileUploadResponse{
		ID: fileUpload.ID,
	}
	return writeJSON(w, http.StatusCreated, resp)
}

// TO DO need to match mapping to either customer, or upload file type
type Mapping struct {
	ExternalID  int
	StartedAt   int
	CancelledAt int
	Amount      int
	Currency    int
	Period      int
}

func ingestSubsFromFile(r io.Reader) ([]data.Subscription, error) {
	// handling csv files only right now
	reader := csv.NewReader(r)
	for {
		// read row by row
		row, err := reader.Read()
		if err != nil {
			// until end of file
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read csv error: %s", err)
		}
		fmt.Println(row)
	}

	return nil, nil
}

func ProcessSubscriptions(subs []data.Subscription) error {
	return nil
}
