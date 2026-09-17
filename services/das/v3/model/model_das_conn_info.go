package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DasConnInfo 连接信息
type DasConnInfo struct {

	// 连接ID
	ConnectionId *string `json:"connection_id,omitempty"`

	// 实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称
	InstanceName *string `json:"instance_name,omitempty"`

	// 数据库来源类型
	NetworkType *string `json:"network_type,omitempty"`

	// 数据库引擎类型
	EngineType *string `json:"engine_type,omitempty"`

	// 数据库版本
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 数据库名称
	DatabaseName *string `json:"database_name,omitempty"`

	// 是否保存密码
	IsSavePassword *bool `json:"is_save_password,omitempty"`

	// IP地址
	IpAddress *string `json:"ip_address,omitempty"`

	// 端口号
	Port *int32 `json:"port,omitempty"`

	// 备注
	Remarks *string `json:"remarks,omitempty"`

	// 实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 连接的创建时间
	CreateAt *int64 `json:"create_at,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// sql记录开关
	SqlRecordFlag *bool `json:"sql_record_flag,omitempty"`

	// 连接类型
	ConnShareType *string `json:"conn_share_type,omitempty"`

	// 共享用户数
	SharedCount *int32 `json:"shared_count,omitempty"`

	// 服务类型
	ServiceType *string `json:"service_type,omitempty"`

	// 共享人名称
	SharedUserName *string `json:"shared_user_name,omitempty"`

	// 共享人ID
	SharedUserId *string `json:"shared_user_id,omitempty"`

	// 共享过期时间
	ExpiredTime *int64 `json:"expired_time,omitempty"`
}

func (o DasConnInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DasConnInfo struct{}"
	}

	return strings.Join([]string{"DasConnInfo", string(data)}, " ")
}
