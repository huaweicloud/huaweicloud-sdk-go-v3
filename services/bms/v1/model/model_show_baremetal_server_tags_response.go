package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowBaremetalServerTagsResponse struct {
	//

	Tags           *[]BaremetalServerTag `json:"tags,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowBaremetalServerTagsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBaremetalServerTagsResponse struct{}"
	}

	return strings.Join([]string{"ShowBaremetalServerTagsResponse", string(data)}, " ")
}
