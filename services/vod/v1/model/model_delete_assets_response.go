package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteAssetsResponse struct {
	DeleteResultArray *[]DeleteResult `json:"delete_result_array,omitempty"`
	HttpStatusCode    int             `json:"-"`
}

func (o DeleteAssetsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAssetsResponse struct{}"
	}

	return strings.Join([]string{"DeleteAssetsResponse", string(data)}, " ")
}
