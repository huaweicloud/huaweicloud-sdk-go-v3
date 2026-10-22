package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiAppBasicInfo AI 应用基本信息
type AiAppBasicInfo struct {

	// AI 应用 ID
	Id *string `json:"id,omitempty"`

	// AI 应用名称
	Name *string `json:"name,omitempty"`
}

func (o AiAppBasicInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiAppBasicInfo struct{}"
	}

	return strings.Join([]string{"AiAppBasicInfo", string(data)}, " ")
}
