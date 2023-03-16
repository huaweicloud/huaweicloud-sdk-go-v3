package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新点位表请求结构体
type UpdateDcPointReqDto struct {

	// 点位名称，允许中、数字、英文大小写、下划线、中划线、#%()*特殊字符
	Name *string `json:"name,omitempty"`

	// 点位采集配置
	CollectionConfig *interface{} `json:"collection_config,omitempty"`

	// 设备id
	DeviceId string `json:"device_id"`

	// 属性，允许中、数字、英文大小写、下划线、中划线
	Property string `json:"property"`

	// 点位数据类型
	DataType *string `json:"data_type,omitempty"`

	ProcessingConfig *ProcessingConfigDto `json:"processing_config,omitempty"`
}

func (o UpdateDcPointReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcPointReqDto struct{}"
	}

	return strings.Join([]string{"UpdateDcPointReqDto", string(data)}, " ")
}
