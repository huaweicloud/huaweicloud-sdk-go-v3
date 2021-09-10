package model

import (
	"encoding/json"

	"strings"
)

type IvsExtentionByNameAndIdRequestBody struct {
	Meta *Meta `json:"meta"`

	Data *IvsExtentionByNameAndIdRequestBodyData `json:"data"`
}

func (o IvsExtentionByNameAndIdRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IvsExtentionByNameAndIdRequestBody struct{}"
	}

	return strings.Join([]string{"IvsExtentionByNameAndIdRequestBody", string(data)}, " ")
}
