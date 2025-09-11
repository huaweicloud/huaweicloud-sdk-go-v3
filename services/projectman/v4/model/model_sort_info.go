package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SortInfo struct {

	// 是否升序
	Asc *bool `json:"asc,omitempty"`

	// 排序字段
	Field *string `json:"field,omitempty"`
}

func (o SortInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SortInfo struct{}"
	}

	return strings.Join([]string{"SortInfo", string(data)}, " ")
}
