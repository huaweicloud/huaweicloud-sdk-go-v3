package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TempContentInfo struct {

	// content_id
	ContentId *int32 `json:"content_id,omitempty" xml:"content_id"`

	// content
	Content *[]Content `json:"content,omitempty" xml:"content"`

	// index
	Index *int32 `json:"index,omitempty" xml:"index"`

	// data
	Data *interface{} `json:"data,omitempty" xml:"data"`

	// data_type
	DataType *int32 `json:"data_type,omitempty" xml:"data_type"`
}

func (o TempContentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TempContentInfo struct{}"
	}

	return strings.Join([]string{"TempContentInfo", string(data)}, " ")
}
