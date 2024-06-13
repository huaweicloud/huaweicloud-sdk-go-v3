package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAllConfigValueByTypeAndKeyResponse Response Object
type ShowAllConfigValueByTypeAndKeyResponse struct {

	// 备注
	Comments *string `json:"comments,omitempty"`

	// 配置键
	ConfigKey *string `json:"config_key,omitempty"`

	// 配置类型
	ConfigType *string `json:"config_type,omitempty"`

	// 配置值
	ConfigValue *string `json:"config_value,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 创建者
	CreateUser *string `json:"create_user,omitempty"`

	// UUID
	Id *string `json:"id,omitempty"`

	// 服务id
	TestServiceId *string `json:"test_service_id,omitempty"`

	// 修改时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 修改者
	UpdateUser     *string `json:"update_user,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowAllConfigValueByTypeAndKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAllConfigValueByTypeAndKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowAllConfigValueByTypeAndKeyResponse", string(data)}, " ")
}
