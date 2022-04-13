package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListInstanceConfigsV2Response struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// 配额列表

	Configs        *[]InstanceConfig `json:"configs,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListInstanceConfigsV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConfigsV2Response struct{}"
	}

	return strings.Join([]string{"ListInstanceConfigsV2Response", string(data)}, " ")
}
