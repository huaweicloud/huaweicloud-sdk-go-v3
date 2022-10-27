package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Configuration struct {

	// 数据。
	Data *interface{} `json:"data,omitempty"`

	// 操作时间。
	OperatedAt *sdktime.SdkTime `json:"operated_at,omitempty"`

	// 操作id。
	OperationId *string `json:"operation_id,omitempty"`

	// 配置类型。
	Type *string `json:"type,omitempty"`

	// 配置是否生效。
	IsActivated *bool `json:"is_activated,omitempty"`
}

func (o Configuration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Configuration struct{}"
	}

	return strings.Join([]string{"Configuration", string(data)}, " ")
}
