package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowScriptDetailResponse Response Object
type ShowScriptDetailResponse struct {

	// 脚本ID。
	Id *string `json:"id,omitempty"`

	// 脚本名称。
	Name *string `json:"name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 脚本类型：POWERSHELL/BAT/SHELL。
	Type *string `json:"type,omitempty"`

	// 脚本版本。
	Version *string `json:"version,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 脚本内容。
	Content        *string `json:"content,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowScriptDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowScriptDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowScriptDetailResponse", string(data)}, " ")
}
