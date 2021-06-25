package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowTakeOverTaskDetailsResponse struct {
	Total *int32 `json:"total,omitempty"`

	TaskId *string `json:"task_id,omitempty"`

	TaskStatus *string `json:"task_status,omitempty"`

	Assets         *[]AssetDetails `json:"assets,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowTakeOverTaskDetailsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTakeOverTaskDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowTakeOverTaskDetailsResponse", string(data)}, " ")
}
