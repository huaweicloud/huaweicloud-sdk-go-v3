package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneShowCatalogRequest struct {
}

func (o KeystoneShowCatalogRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneShowCatalogRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowCatalogRequest", string(data)}, " ")
}
