package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDsResponse Response Object
type CreateDsResponse struct {

	// 采集数据源id，节点下唯一
	DsId *string `json:"ds_id,omitempty"`

	// 数据源的连接及采集信息
	Config *interface{} `json:"config,omitempty"`

	// 采集数据源名称，允许中、数字、英文大小写、下划线、中划线
	Name *string `json:"name,omitempty"`

	// 边缘节点id
	EdgeNodeId *string `json:"edge_node_id,omitempty"`

	// 模块id
	ModuleId *string `json:"module_id,omitempty"`

	// 模板id，节点下唯一
	TplId *string `json:"tpl_id,omitempty"`

	// 质量上报开关，不携带或值不为true，默认为false
	QualityReport *bool `json:"quality_report,omitempty"`

	// 应用ID
	EdgeAppName *string `json:"edge_app_name,omitempty"`

	// 数采连接信息
	ConnectionInfo *interface{} `json:"connection_info,omitempty"`

	// 数采连接状态,stopped|running
	ModuleState *string `json:"module_state,omitempty"`

	// 数采连接下点位数
	Count *int64 `json:"count,omitempty"`

	// 创建时间
	CreateTime *string `json:"create_time,omitempty"`

	// 最后一次修改时间
	UpdateTime *string `json:"update_time,omitempty"`

	// 数采配置是否已同步，已同步：true,未同步：false
	Synchronized *bool `json:"synchronized,omitempty"`

	// 数采配置同步时间
	SynchronizedTime *string `json:"synchronized_time,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o CreateDsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDsResponse struct{}"
	}

	return strings.Join([]string{"CreateDsResponse", string(data)}, " ")
}
