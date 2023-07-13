package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateEdgeModuleReqDto 更新边缘模块请求结构体
type UpdateEdgeModuleReqDto struct {

	// 边缘应用版本
	AppVersion *string `json:"app_version,omitempty"`

	// 边缘模块名称
	ModuleName *string `json:"module_name,omitempty"`

	ContainerSettings *ContainerSettingsReqDto `json:"container_settings,omitempty"`
}

func (o UpdateEdgeModuleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeModuleReqDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeModuleReqDto", string(data)}, " ")
}
