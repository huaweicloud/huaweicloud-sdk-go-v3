package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建点位表配置请求结构体
type CreateDcPointReqDto struct {

	// 点位表id，数据源下唯一
	PointId string `json:"point_id"`

	// 点位名称，允许中、数字、英文大小写、下划线、中划线、#%()*特殊字符
	Name string `json:"name"`

	// 点位数据类型
	DataType *string `json:"data_type,omitempty"`

	// 点位采集配置
	CollectionConfig *interface{} `json:"collection_config"`

	// 设备id
	DeviceId string `json:"device_id"`

	// 属性，允许中、数字、英文大小写、下划线、中划线
	Property string `json:"property"`

	ProcessingConfig *ProcessingConfigDto `json:"processing_config,omitempty"`
}

func (o CreateDcPointReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDcPointReqDto struct{}"
	}

	return strings.Join([]string{"CreateDcPointReqDto", string(data)}, " ")
}
