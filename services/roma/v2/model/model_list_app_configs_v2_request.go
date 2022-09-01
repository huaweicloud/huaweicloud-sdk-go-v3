package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListAppConfigsV2Request struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 偏移量，表示从此偏移量开始查询，偏移量小于0时，自动转换为0
	Offset *int64 `json:"offset,omitempty" xml:"offset"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 应用编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 应用配置名称
	ConfigName *string `json:"config_name,omitempty" xml:"config_name"`

	// 应用名称
	RomaAppName *string `json:"roma_app_name,omitempty" xml:"roma_app_name"`
}

func (o ListAppConfigsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppConfigsV2Request struct{}"
	}

	return strings.Join([]string{"ListAppConfigsV2Request", string(data)}, " ")
}
