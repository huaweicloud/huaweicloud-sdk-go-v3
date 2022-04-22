package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建边缘节点请求结构体。
type EdgeNodeCreation struct {

	// 边缘节点名称，只允许中、数字、英文大小写、中划线、下划线
	Name string `json:"name"`

	// 节点所属资源类型：advanced|standard
	Type string `json:"type"`

	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"instance_id,omitempty"`

	// 资源空间id，对应IOTDA云服务接口参数中的app_id。
	SpaceId *string `json:"space_id,omitempty"`

	// 资源id列表，创建节点时需绑定已购买的资源包，可以叠加节点功能。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 节点的安全等级，MEDIUM边缘节数据上报不进行加密，HIGH对数据上报进行加密。
	SecurityLevel *string `json:"security_level,omitempty"`

	// 节点的存储周期，默认0天，取值范围0~7天，0天则不存储。
	StoragePeriod *int32 `json:"storage_period,omitempty"`

	// 华为AI加速卡类型，如NPU、GPU
	AiCardType *string `json:"ai_card_type,omitempty"`

	// 边缘节点在IEF日志配置参数
	LogConfigs *[]LogConfigDto `json:"log_configs,omitempty"`

	// 用户预置第三方边缘应用
	Apps *[]EdgeAppInstanceDto `json:"apps,omitempty"`
}

func (o EdgeNodeCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeCreation struct{}"
	}

	return strings.Join([]string{"EdgeNodeCreation", string(data)}, " ")
}
