package model

import (
	"encoding/json"

	"strings"
)

type ShowTaskResqTaskinfoContent1 struct {
	// content_type

	ContentType *int32 `json:"content_type,omitempty"`

	Content *ShowTaskResqTaskinfoContent `json:"content,omitempty"`
}

func (o ShowTaskResqTaskinfoContent1) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskResqTaskinfoContent1 struct{}"
	}

	return strings.Join([]string{"ShowTaskResqTaskinfoContent1", string(data)}, " ")
}
