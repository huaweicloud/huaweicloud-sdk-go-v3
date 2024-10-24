package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelVersionInfo Model Version详情信息
type ModelVersionInfo struct {

	// 模型版本ID，32~36位的英文、数字、短横组合，系统自动生成无法修改，输入不生效。
	Id *string `json:"id,omitempty"`

	// 模型版本名称, 只能包含中文、字母、数字、下划线、中划线、点、空格
	Name string `json:"name"`

	// cap白名单
	CapWhiteList *[]string `json:"cap_white_list,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	Config *ModelConfig `json:"config"`

	// 创建时间
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	CreateUser *User `json:"create_user,omitempty"`
}

func (o ModelVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelVersionInfo struct{}"
	}

	return strings.Join([]string{"ModelVersionInfo", string(data)}, " ")
}
