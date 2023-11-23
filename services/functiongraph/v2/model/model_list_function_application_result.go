package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListFunctionApplicationResult struct {

	// 应用id
	Id *string `json:"id,omitempty"`

	// 应用名称
	Name *string `json:"name,omitempty"`

	// 应用状态
	Status *string `json:"status,omitempty"`

	// 最后修改时间
	LastModifiedTime *int64 `json:"last_modified_time,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`
}

func (o ListFunctionApplicationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionApplicationResult struct{}"
	}

	return strings.Join([]string{"ListFunctionApplicationResult", string(data)}, " ")
}
