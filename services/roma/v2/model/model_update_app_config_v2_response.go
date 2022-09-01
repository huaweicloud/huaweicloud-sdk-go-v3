package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateAppConfigV2Response struct {

	// 应用配置编号
	Id *string `json:"id,omitempty" xml:"id"`

	// 应用编号
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 应用配置类型
	ConfigType *string `json:"config_type,omitempty" xml:"config_type"`

	// 应用配置名称
	ConfigName *string `json:"config_name,omitempty" xml:"config_name"`

	// 应用配置值
	ConfigValue *string `json:"config_value,omitempty" xml:"config_value"`

	// 应用配置更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty" xml:"update_time"`

	// 应用配置描述
	Description    *string `json:"description,omitempty" xml:"description"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateAppConfigV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAppConfigV2Response struct{}"
	}

	return strings.Join([]string{"UpdateAppConfigV2Response", string(data)}, " ")
}
