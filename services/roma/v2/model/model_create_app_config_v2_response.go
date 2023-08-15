package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppConfigV2Response Response Object
type CreateAppConfigV2Response struct {

	// 应用配置编号
	Id *string `json:"id,omitempty"`

	// 应用编号
	AppId *string `json:"app_id,omitempty"`

	// 应用配置类型
	ConfigType *string `json:"config_type,omitempty"`

	// 应用配置名称
	ConfigName *string `json:"config_name,omitempty"`

	// 应用配置值
	ConfigValue *string `json:"config_value,omitempty"`

	// 应用配置更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 应用配置描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAppConfigV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppConfigV2Response struct{}"
	}

	return strings.Join([]string{"CreateAppConfigV2Response", string(data)}, " ")
}
