package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiListConnectionsInfoRespDasConnInfoList struct {

	// 连接Id
	ConnectionId *string `json:"connection_id,omitempty"`

	// 实例Id
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 数据库来源
	NetworkType *string `json:"network_type,omitempty"`

	// 数据库引擎
	EngineType *string `json:"engine_type,omitempty"`

	// 数据存储版本
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 是否保存密码
	IsSavePassword *bool `json:"is_save_password,omitempty"`

	// ip地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 端口号
	Port *int32 `json:"port,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 连接的创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 状态，NORMAL-正常，INSTANCE_DELETE-实例删除
	Status *string `json:"status,omitempty"`

	// 连接类型，NORMAL-正常连接，SHARE-共享连接
	ConnShareType *string `json:"conn_share_type,omitempty"`

	// 共享人名称
	SharedUserName *string `json:"shared_user_name,omitempty"`

	// 共享人ID
	SharedUserId *string `json:"shared_user_id,omitempty"`

	// 共享过期时间
	ExpiredTime *int64 `json:"expired_time,omitempty"`
}

func (o ApiListConnectionsInfoRespDasConnInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiListConnectionsInfoRespDasConnInfoList struct{}"
	}

	return strings.Join([]string{"ApiListConnectionsInfoRespDasConnInfoList", string(data)}, " ")
}
