package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Contents struct {

	// content_id
	ContentId *int32 `json:"content_id,omitempty" xml:"content_id"`

	// content
	Content *[]Content `json:"content,omitempty" xml:"content"`

	// index
	Index *int32 `json:"index,omitempty" xml:"index"`

	// selected_temp_name
	SelectedTempName *string `json:"selected_temp_name,omitempty" xml:"selected_temp_name"`

	// data
	Data *interface{} `json:"data,omitempty" xml:"data"`

	// data_type
	DataType *int32 `json:"data_type,omitempty" xml:"data_type"`
}

func (o Contents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Contents struct{}"
	}

	return strings.Join([]string{"Contents", string(data)}, " ")
}
