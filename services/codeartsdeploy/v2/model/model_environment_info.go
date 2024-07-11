package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnvironmentInfo 环境信息
type EnvironmentInfo struct {

	// 环境id
	Id *string `json:"id,omitempty"`

	// 环境名称
	Name *string `json:"name,omitempty"`

	// 应用id
	AppId *string `json:"app_id,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 项目id
	ProjectId *string `json:"project_id,omitempty"`
}

func (o EnvironmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvironmentInfo struct{}"
	}

	return strings.Join([]string{"EnvironmentInfo", string(data)}, " ")
}
