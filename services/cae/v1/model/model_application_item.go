package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplicationItem 应用信息。
type ApplicationItem struct {

	// 应用ID。
	Id *string `json:"id,omitempty"`

	// 应用名称。
	Name *string `json:"name,omitempty"`

	// 应用附加属性，当前只支持description参数。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 创建时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ApplicationItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicationItem struct{}"
	}

	return strings.Join([]string{"ApplicationItem", string(data)}, " ")
}
