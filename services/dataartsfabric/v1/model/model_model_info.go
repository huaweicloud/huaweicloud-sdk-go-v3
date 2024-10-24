package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelInfo 模型详情信息
type ModelInfo struct {

	// 可见性：  - PRIVATE：私有  - PUBLIC：公共  默认为PRIVATE
	Visibility *string `json:"visibility,omitempty"`

	// 模型ID，32~36位的英文、数字、短横组合
	Id *string `json:"id,omitempty"`

	// 一个Model的名称，只能包含中文、字母、数字、下划线、中划线、点、空格
	Name *string `json:"name,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Type *ModelType `json:"type,omitempty"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	CurrentVersion *ModelVersionInfo `json:"current_version,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`

	UpdateUser *User `json:"update_user,omitempty"`
}

func (o ModelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelInfo struct{}"
	}

	return strings.Join([]string{"ModelInfo", string(data)}, " ")
}
