package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplaySourceConfigResp 流量回放源数据库配置返回体
type ReplaySourceConfigResp struct {

	// 流量回放源数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 流量回放源数据库类型
	SourceEndpointType *string `json:"source_endpoint_type,omitempty"`

	// 流量回放源数据库来源
	SourceDbOrigin *string `json:"source_db_origin,omitempty"`

	// 流量回放流量文件来源
	TrafficSource *string `json:"traffic_source,omitempty"`

	// 流量回放流量类型
	SupportTrafficTypes *[]string `json:"support_traffic_types,omitempty"`

	// 流量回放agent支持的操作系统
	SupportOsTypes *[]string `json:"support_os_types,omitempty"`

	// 流量回放引擎支持的镜像库
	SupportMirrorDatabaseTypes *[]string `json:"support_mirror_database_types,omitempty"`
}

func (o ReplaySourceConfigResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplaySourceConfigResp struct{}"
	}

	return strings.Join([]string{"ReplaySourceConfigResp", string(data)}, " ")
}
