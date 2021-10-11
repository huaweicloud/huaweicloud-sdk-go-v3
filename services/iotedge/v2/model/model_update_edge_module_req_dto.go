package model

import (
	"encoding/json"

	"strings"
)

// 创建边缘模块请求结构体
type UpdateEdgeModuleReqDto struct {
	// 边缘应用版本

	AppVersion string `json:"app_version"`
}

func (o UpdateEdgeModuleReqDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateEdgeModuleReqDto struct{}"
	}

	return strings.Join([]string{"UpdateEdgeModuleReqDto", string(data)}, " ")
}
