package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportAppRsp 导入应用响应体
type ImportAppRsp struct {

	// 源应用id
	SourceAppId *string `json:"source_app_id,omitempty"`

	// 目标应用id
	DestinationAppId *string `json:"destination_app_id,omitempty"`

	// 目标应用名称
	DestinationAppName *string `json:"destination_app_name,omitempty"`

	// 应用版本
	Version *string `json:"version,omitempty"`

	// 导入结果信息，仅在导入失败时会返回
	Message *string `json:"message,omitempty"`

	// 导入结果状态
	Status *string `json:"status,omitempty"`
}

func (o ImportAppRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportAppRsp struct{}"
	}

	return strings.Join([]string{"ImportAppRsp", string(data)}, " ")
}
