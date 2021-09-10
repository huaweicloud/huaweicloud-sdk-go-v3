package model

import (
	"encoding/json"

	"strings"
)

type IvsExtentionByIdCardImageRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsExtentionByIdCardImageRequestBodyData `json:"data"`
}

func (o IvsExtentionByIdCardImageRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IvsExtentionByIdCardImageRequestBody struct{}"
	}

	return strings.Join([]string{"IvsExtentionByIdCardImageRequestBody", string(data)}, " ")
}
