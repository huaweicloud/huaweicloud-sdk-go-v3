package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EdgeNodeCreation 创建边缘节点请求结构体。
type EdgeNodeCreation struct {

	// 边缘节点ID
	EdgeNodeId *string `json:"edge_node_id,omitempty"`

	// 边缘节点名称，只允许中、数字、英文大小写、中划线、下划线
	Name string `json:"name"`

	// 边缘节点类型：lite|advanced|standard。lite表示基础版边缘节点，advanced或standard表示专业版边缘节点。
	Type string `json:"type"`

	// 边缘节点注册使用的验证码，如果不输入则平台随机生成。
	VerifyCode *string `json:"verify_code,omitempty"`

	// 验证码的有效时间单位秒，默认1800秒，范围为1~864000，过期后平台会随机生成。
	TimeOut *int32 `json:"time_out,omitempty"`

	// 边缘节点系统架构。包括：arm64，arm32，x86_64。
	Arch *string `json:"arch,omitempty"`

	// 边缘节点系统类型。包括：generalLinux通用系统，openHarmony鸿蒙系统。
	OsType *string `json:"os_type,omitempty"`

	// 实例ID。物理多租下各实例的唯一标识，一般华为云租户无需携带该参数，仅在物理多租场景下从管理面访问API时需要携带该参数。
	InstanceId *string `json:"instance_id,omitempty"`

	// 资源空间id，对应IOTDA云服务接口参数中的app_id。
	SpaceId *string `json:"space_id,omitempty"`

	// 资源id列表，创建节点时需绑定已购买的资源包，资源可叠加。
	ResourceIds *[]string `json:"resource_ids,omitempty"`

	// 节点的安全等级，MEDIUM表示本地明文存储，HIGH表示本地加密存储。
	SecurityLevel *string `json:"security_level,omitempty"`

	// 节点的可靠性等级，LOW表示中级别，MEDIUM表示高级别。详细功能请参考“用户指南>管理边缘节点>注册节点”。
	ReliabilityLevel *string `json:"reliability_level,omitempty"`

	// 节点的存储周期，默认0天，取值范围0~7天，0天则不存储。
	StoragePeriod *int32 `json:"storage_period,omitempty"`

	// AI加速卡类型，如昇腾AI加速卡NPU、图像处理加速卡GPU。
	AiCardType *string `json:"ai_card_type,omitempty"`

	// npu驱动动态库路径
	NpuLibraryPath *string `json:"npu_library_path,omitempty"`

	BasePath *BasePathDto `json:"base_path,omitempty"`

	// 边缘节点在IEF日志配置参数，仅专业版支持。
	LogConfigs *[]LogConfigDto `json:"log_configs,omitempty"`

	// 需要自动安装的边缘应用。此处可填写控制台“应用管理”页面中列出的业务应用与驱动应用。
	Apps *[]EdgeAppInstanceDto `json:"apps,omitempty"`

	// 网络接入方式类型
	NetworkAccessPoint *string `json:"network_access_point,omitempty"`

	// 网关型号
	HardwareModel *string `json:"hardware_model,omitempty"`

	OfflineCacheConfigs *OfflineCacheConfigsDto `json:"offline_cache_configs,omitempty"`

	DeviceAuthInfo *DeviceAuthInfoDto `json:"device_auth_info,omitempty"`

	// 节点使用的数据格式，默认为iotda物模型1.0格式，可以选择属性平铺数据格式flat_json
	DeviceDataFormat *string `json:"device_data_format,omitempty"`

	// 自动升级系统应用的节点开关，默认为关闭：OFF，IMMEDIATE表示节点开关打开
	AutomaticUpgrade *string `json:"automatic_upgrade,omitempty"`

	DeviceDataRecord *DeviceDataRecord `json:"device_data_record,omitempty"`

	// omagent监控运维工具是否上报指标
	MetricReport *string `json:"metric_report,omitempty"`

	// iotda的南向接入地址
	IotdaSouthAccess *string `json:"iotda_south_access,omitempty"`
}

func (o EdgeNodeCreation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EdgeNodeCreation struct{}"
	}

	return strings.Join([]string{"EdgeNodeCreation", string(data)}, " ")
}
