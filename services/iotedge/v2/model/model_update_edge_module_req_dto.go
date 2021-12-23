package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建边缘模块请求结构体
type UpdateEdgeModuleReqDto struct {
	// 边缘应用版本

	AppVersion string `json:"app_version"`
}

func (o UpdateEdgeModuleReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeModuleReqDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeModuleReqDto", string(data)}, " ")
}
