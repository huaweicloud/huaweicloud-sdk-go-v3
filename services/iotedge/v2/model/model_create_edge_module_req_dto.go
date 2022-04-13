package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建边缘模块请求结构体
type CreateEdgeModuleReqDto struct {
	// 边缘应用名称

	EdgeAppId string `json:"edge_app_id"`
	// 边缘应用版本

	AppVersion string `json:"app_version"`
}

func (o CreateEdgeModuleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeModuleReqDto struct{}"
	}

	return strings.Join([]string{"CreateEdgeModuleReqDto", string(data)}, " ")
}
