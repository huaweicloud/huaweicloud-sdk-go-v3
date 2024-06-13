package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNodeReqDto 修改边缘节点请求结构体。
type UpdateNodeReqDto struct {

	// 边缘节点名称，只允许中、数字、英文大小写、中划线、下划线
	Name *string `json:"name,omitempty"`

	// 节点的存储周期，默认0天，取值范围0~7天，0天则不存储。
	StoragePeriod *int32 `json:"storage_period,omitempty"`

	// 边缘节点在IEF日志配置参数
	LogConfigs *[]LogConfigDto `json:"log_configs,omitempty"`

	HaConfig *HaConfigDto `json:"ha_config,omitempty"`

	// 网关型号
	HardwareModel *string `json:"hardware_model,omitempty"`

	// npu驱动动态库路径
	NpuLibraryPath *string `json:"npu_library_path,omitempty"`

	// 自动升级系统应用的节点开关，默认为关闭：OFF，IMMEDIATE表示节点开关打开
	AutomaticUpgrade *string `json:"automatic_upgrade,omitempty"`

	DeviceDataRecord *DeviceDataRecord `json:"device_data_record,omitempty"`

	// omagent监控运维工具是否上报指标
	MetricReport *string `json:"metric_report,omitempty"`

	OfflineCacheConfigs *UpdateOfflineCacheConfigsDto `json:"offline_cache_configs,omitempty"`
}

func (o UpdateNodeReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNodeReqDto struct{}"
	}

	return strings.Join([]string{"UpdateNodeReqDto", string(data)}, " ")
}
