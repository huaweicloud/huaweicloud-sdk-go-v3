package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TransactionConfigItem URL跟踪视图配置。
type TransactionConfigItem struct {

	// 配置id。
	Id *int64 `json:"id,omitempty"`

	// 应用id。
	BusinessId *int64 `json:"business_id,omitempty"`

	// 环境id。
	EnvId *int64 `json:"env_id,omitempty"`

	// 请求方式。
	Method *string `json:"method,omitempty"`

	// 环境名称。
	EnvName *string `json:"env_name,omitempty"`

	// region显示英文名称。
	Region *string `json:"region,omitempty"`

	// 类型。
	Type *string `json:"type,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// url地址。
	Url *string `json:"url,omitempty"`
}

func (o TransactionConfigItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransactionConfigItem struct{}"
	}

	return strings.Join([]string{"TransactionConfigItem", string(data)}, " ")
}
