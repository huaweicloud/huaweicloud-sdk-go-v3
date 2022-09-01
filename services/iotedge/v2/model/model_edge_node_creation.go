package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建边缘节点请求结构体。
type EdgeNodeCreation struct {

	// 边缘节点ID
	EdgeNodeId *string `json:"edge_node_id,omitempty" xml:"edge_node_id"`

	// 边缘节点名称，只允许中、数字、英文大小写、中划线、下划线
	Name string `json:"name" xml:"name"`

	// 节点所属资源类型：advanced|standard
	Type string `json:"type" xml:"type"`

	// 边缘节点注册使用的验证码，如果不输入则平台随机生成。
	VerifyCode *string `json:"verify_code,omitempty" xml:"verify_code"`

	// 验证码的有效时间单位秒，默认1800秒，范围为1~864000，过期后平台会随机生成。
	TimeOut *int32 `json:"time_out,omitempty" xml:"time_out"`

	// 系统架构。包括：arm64，arm32，x86_64。
	Arch *string `json:"arch,omitempty" xml:"arch"`

	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 资源空间id，对应IOTDA云服务接口参数中的app_id。
	SpaceId *string `json:"space_id,omitempty" xml:"space_id"`

	// 资源id列表，创建节点时需绑定已购买的资源包，资源可叠加。
	ResourceIds *[]string `json:"resource_ids,omitempty" xml:"resource_ids"`

	// 节点的安全等级，MEDIUM表示本地明文存储，HIGH表示本地加密存储。
	SecurityLevel *string `json:"security_level,omitempty" xml:"security_level"`

	// 节点的存储周期，默认0天，取值范围0~7天，0天则不存储。
	StoragePeriod *int32 `json:"storage_period,omitempty" xml:"storage_period"`

	// 华为AI加速卡类型，如NPU、GPU。
	AiCardType *string `json:"ai_card_type,omitempty" xml:"ai_card_type"`

	BasePath *BasePathDto `json:"base_path,omitempty" xml:"base_path"`

	// 边缘节点在IEF日志配置参数，仅高级版支持。
	LogConfigs *[]LogConfigDto `json:"log_configs,omitempty" xml:"log_configs"`

	// 用户预置第三方边缘应用
	Apps *[]EdgeAppInstanceDto `json:"apps,omitempty" xml:"apps"`

	// 网关型号
	HardwareModel *string `json:"hardware_model,omitempty" xml:"hardware_model"`
}

func (o EdgeNodeCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeCreation struct{}"
	}

	return strings.Join([]string{"EdgeNodeCreation", string(data)}, " ")
}
