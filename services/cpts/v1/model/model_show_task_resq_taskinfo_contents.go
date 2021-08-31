package model

import (
	"encoding/json"

	"strings"
)

type ShowTaskResqTaskinfoContents struct {
	// content_id

	ContentId *int32 `json:"content_id,omitempty"`
	// content

	Content *[]ShowTaskResqTaskinfoContent1 `json:"content,omitempty"`
	// index

	Index *int32 `json:"index,omitempty"`
	// selected_temp_name

	SelectedTempName *string `json:"selected_temp_name,omitempty"`
	// data

	Data *string `json:"data,omitempty"`
	// data_type

	DataType *int32 `json:"data_type,omitempty"`
}

func (o ShowTaskResqTaskinfoContents) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowTaskResqTaskinfoContents struct{}"
	}

	return strings.Join([]string{"ShowTaskResqTaskinfoContents", string(data)}, " ")
}
