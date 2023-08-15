package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IndicatorInfo struct {

	// 监控指标名称。
	IndicatorName *string `json:"indicator_name,omitempty"`

	// 采集模块名称。
	PluginName *string `json:"plugin_name,omitempty"`

	// 默认采集频率。
	DefaultCollectRate *string `json:"default_collect_rate,omitempty"`

	// 支持的集群版本。
	SupportDatastoreVersion *string `json:"support_datastore_version,omitempty"`
}

func (o IndicatorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IndicatorInfo struct{}"
	}

	return strings.Join([]string{"IndicatorInfo", string(data)}, " ")
}
