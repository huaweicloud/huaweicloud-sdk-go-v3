package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TempDetailInfo struct {

	// description
	Description *string `json:"description,omitempty" xml:"description"`

	// id
	Id *int32 `json:"id,omitempty" xml:"id"`

	// 是否被引用
	IsQuoted *bool `json:"is_quoted,omitempty" xml:"is_quoted"`

	// name
	Name *string `json:"name,omitempty" xml:"name"`

	// temp_type
	TempType *int32 `json:"temp_type,omitempty" xml:"temp_type"`

	// update_time
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time"`
}

func (o TempDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TempDetailInfo struct{}"
	}

	return strings.Join([]string{"TempDetailInfo", string(data)}, " ")
}
