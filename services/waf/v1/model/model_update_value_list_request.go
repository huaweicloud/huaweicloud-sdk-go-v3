package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateValueListRequest struct {
	// 引用表id

	Valuelistid string `json:"valuelistid"`

	Body *UpdateValueListRequestBody `json:"body,omitempty"`
}

func (o UpdateValueListRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateValueListRequest struct{}"
	}

	return strings.Join([]string{"UpdateValueListRequest", string(data)}, " ")
}
