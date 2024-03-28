package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DatasourceConnection 查询经典型跨源连接的响应参数。
type DatasourceConnection struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息为空。
	Message *string `json:"message,omitempty"`

	// 连接ID，用于标识跨源连接的UUID。
	ConnectionId *string `json:"connection_id,omitempty"`

	// 创建连接时，用户填写的队列的访问地址。
	Destination *string `json:"destination,omitempty"`

	// 连接状态。CREATING：跨源连接正在创建中；ACTIVE：跨源连接创建成功，与目的地址连接正常；FAILED：跨源连接创建失败；DELETED：跨源连接已被删除。
	State *string `json:"state,omitempty"`

	// 正在创建的跨源连接进度，显示0.0至1.0代表0%至100%。
	Process *float64 `json:"process,omitempty"`

	// 创建连接时，用户自定义的连接名称。
	Name *string `json:"name,omitempty"`

	// 用于建立跨源关联表时，需要使用到的连接url。
	ConnectionUrl *string `json:"connection_url,omitempty"`

	// Serverless Spark队列名称。SQL队列模式下建立的跨源连接，该字段为空。
	ClusterName *string `json:"cluster_name,omitempty"`

	// 创建连接时，用户指定的对端服务（CloudTable/CloudTable.OpenTSDB/MRS.OpenTSDB/DWS/RDS/CSS）。
	Service *string `json:"service,omitempty"`

	// 创建连接的时间。为UTC的时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 当前跨源已绑定的队列名
	QueueList *[]string `json:"queue_list,omitempty"`
}

func (o DatasourceConnection) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DatasourceConnection struct{}"
	}

	return strings.Join([]string{"DatasourceConnection", string(data)}, " ")
}
