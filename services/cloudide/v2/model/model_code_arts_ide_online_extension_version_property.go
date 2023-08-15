package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CodeArtsIdeOnlineExtensionVersionProperty struct {

	// id
	Id *int32 `json:"id,omitempty"`

	// 参数名
	PropertyName *string `json:"property_name,omitempty"`

	// 参数值
	PropertyValue *string `json:"property_value,omitempty"`

	// 插件版本id
	ExtensionVersionId *string `json:"extension_version_id,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o CodeArtsIdeOnlineExtensionVersionProperty) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CodeArtsIdeOnlineExtensionVersionProperty struct{}"
	}

	return strings.Join([]string{"CodeArtsIdeOnlineExtensionVersionProperty", string(data)}, " ")
}
