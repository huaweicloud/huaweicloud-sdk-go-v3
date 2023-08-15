package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateEdgeModuleReqDto 创建边缘模块请求结构体
type CreateEdgeModuleReqDto struct {

	// 边缘应用名称
	EdgeAppId string `json:"edge_app_id"`

	// 边缘应用版本
	AppVersion string `json:"app_version"`

	// 边缘模块名称
	ModuleName *string `json:"module_name,omitempty"`

	ContainerSettings *ContainerSettingsReqDto `json:"container_settings,omitempty"`
}

func (o CreateEdgeModuleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeModuleReqDto struct{}"
	}

	return strings.Join([]string{"CreateEdgeModuleReqDto", string(data)}, " ")
}
