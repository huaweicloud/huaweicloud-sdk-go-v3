package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceDetail struct {

	// 实例ID。
	Id *string `json:"id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 实例别名。
	Alias *string `json:"alias,omitempty"`

	// 项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 集群模式。
	ClusterMode *string `json:"cluster_mode,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// bpdomain_id
	BpdomainId *string `json:"bpdomain_id,omitempty"`

	// 账户ID。
	UserId *string `json:"user_id,omitempty"`

	// 数据库版本。
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// 数据库类型。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 创建时间。
	CreateAt *string `json:"create_at,omitempty"`

	// 更新时间。
	UpdateAt *string `json:"update_at,omitempty"`

	// 删除时间。
	DeleteAt *string `json:"delete_at,omitempty"`

	// 是否有版本可升级。
	NewVersionAvailable *bool `json:"new_version_available,omitempty"`

	// 是否有版本可回滚。
	RollbackVersionAvailable *bool `json:"rollback_version_available,omitempty"`

	// 是否有版本可降级。
	DegradeVersionAvailable *bool `json:"degrade_version_available,omitempty"`

	// 公共ip。
	PublicIp *string `json:"public_ip,omitempty"`

	// 端口。
	Port *string `json:"port,omitempty"`

	// 创建失败原因编码。
	CreateFailErrorCode *string `json:"create_fail_error_code,omitempty"`

	// 时区。
	TimeZone *string `json:"time_zone,omitempty"`

	// 付费模式。
	PayModel *string `json:"pay_model,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 周期。
	Period *int32 `json:"period,omitempty"`

	// 是否冻结。
	IsFrozen *bool `json:"is_frozen,omitempty"`

	// 冻结时间。
	FrozenTime *string `json:"frozen_time,omitempty"`

	// 锁状态。
	Actions *[]ActionInfo `json:"actions,omitempty"`

	// 是否只有默认组。
	OnlyDefaultGroup *bool `json:"only_default_group,omitempty"`

	// 组信息。
	Groups *[]DdmGroupInfo `json:"groups,omitempty"`

	// 其他信息。
	ExtendMap map[string]string `json:"extend_map,omitempty"`

	// 标签信息。
	TagsInfo *[]TagsInfo `json:"tags_info,omitempty"`

	// 管理员账号。
	AdminUserName *string `json:"admin_user_name,omitempty"`

	// 绑定eip信息。
	EipBindingInfo *interface{} `json:"eip_binding_info,omitempty"`

	// 是否支持ssl。
	EnableSsl *int32 `json:"enable_ssl,omitempty"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}
