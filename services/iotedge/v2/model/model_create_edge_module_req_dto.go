package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建边缘模块请求结构体
type CreateEdgeModuleReqDto struct {

	// 边缘应用名称
	EdgeAppId string `json:"edge_app_id" xml:"edge_app_id"`

	// 边缘应用版本
	AppVersion string `json:"app_version" xml:"app_version"`

	// 边缘模块名称
	ModuleName *string `json:"module_name,omitempty" xml:"module_name"`

	ContainerSettings *ContainerSettingsReqDto `json:"container_settings,omitempty" xml:"container_settings"`
}

func (o CreateEdgeModuleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeModuleReqDto struct{}"
	}

	return strings.Join([]string{"CreateEdgeModuleReqDto", string(data)}, " ")
}
