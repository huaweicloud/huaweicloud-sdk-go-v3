package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateDcPointResponse struct {

	// 点位表id，数据源下唯一
	PointId *string `json:"point_id,omitempty"`

	// 点位名称，允许中、数字、英文大小写、下划线、中划线、#%()*特殊字符
	Name *string `json:"name,omitempty"`

	// 点位采集配置
	CollectionConfig *interface{} `json:"collection_config,omitempty"`

	// 设备id
	DeviceId *string `json:"device_id,omitempty"`

	// 属性，允许中、数字、英文大小写、下划线、中划线
	Property *string `json:"property,omitempty"`

	// 点位数据类型
	DataType *string `json:"data_type,omitempty"`

	// 采集数据源id，节点下唯一
	DsId *string `json:"ds_id,omitempty"`

	ProcessingConfig *ProcessingConfigDto `json:"processing_config,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateDcPointResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDcPointResponse struct{}"
	}

	return strings.Join([]string{"UpdateDcPointResponse", string(data)}, " ")
}
