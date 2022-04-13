package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TempContentInfo struct {
	// content_id

	ContentId *int32 `json:"content_id,omitempty"`
	// content

	Content *[]Content `json:"content,omitempty"`
	// index

	Index *int32 `json:"index,omitempty"`
	// data

	Data *string `json:"data,omitempty"`
	// data_type

	DataType *int32 `json:"data_type,omitempty"`
}

func (o TempContentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TempContentInfo struct{}"
	}

	return strings.Join([]string{"TempContentInfo", string(data)}, " ")
}
