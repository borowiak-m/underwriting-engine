package handler

import (
	"io"
	"net/http"

	"github.com/borowiak-m/underwriting-engine/data"
)

func Upload(w http.ResponseWriter, r *http.Request) error {
	// TO DO QC checks
	subs, err := ingestSubsFromFile(r.Body)
	if err != nil {
		return err
	}
	ProcessSubscriptions(subs)
	return nil
}

func ingestSubsFromFile(r io.Reader) ([]data.Subscription, error)

func ProcessSubscriptions(subs []data.Subscription) error
