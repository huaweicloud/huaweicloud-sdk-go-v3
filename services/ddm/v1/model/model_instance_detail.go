package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceDetail struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **参数范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。
	Id *string `json:"id,omitempty"`

	// **参数解释**：  实例名称。  **参数范围**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  实例别名。  **参数范围**：  不涉及。
	Alias *string `json:"alias,omitempty"`

	// **参数解释**：  租户在某一Region下的project ID。  获取方法请参见[获取项目ID](https://support.huaweicloud.com/api-ddm/ddm_api_01_0063.html)。  **参数范围**：  只能由英文字母、数字组成，且长度为32个字符。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：  集群模式。  **参数范围**：  不涉及。
	ClusterMode *string `json:"cluster_mode,omitempty"`

	// **参数解释**：  状态。  **参数范围**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  bpdomain_id。  **参数范围**：  不涉及。
	BpdomainId *string `json:"bpdomain_id,omitempty"`

	// **参数解释**：  账户ID。  **参数范围**：  不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：  数据库版本。  **参数范围**：  不涉及。
	DatastoreVersion *string `json:"datastore_version,omitempty"`

	// **参数解释**：  数据库类型。  **参数范围**：  不涉及。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// **参数解释**：  创建时间。  **参数范围**：  不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释**：  更新时间。  **参数范围**：  不涉及。
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释**：  删除时间。  **参数范围**：  不涉及。
	DeleteAt *string `json:"delete_at,omitempty"`

	// **参数解释**：  是否有版本可升级。  **参数范围**：  不涉及。
	NewVersionAvailable *bool `json:"new_version_available,omitempty"`

	// **参数解释**：  是否有版本可回滚。  **参数范围**：  不涉及。
	RollbackVersionAvailable *bool `json:"rollback_version_available,omitempty"`

	// **参数解释**：  是否有版本可降级。  **参数范围**：  不涉及。
	DegradeVersionAvailable *bool `json:"degrade_version_available,omitempty"`

	// **参数解释**：  公共ip。  **参数范围**：  不涉及。
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**：  端口。  **参数范围**：  不涉及。
	Port *string `json:"port,omitempty"`

	// **参数解释**：  创建失败原因编码。  **参数范围**：  不涉及。
	CreateFailErrorCode *string `json:"create_fail_error_code,omitempty"`

	// **参数解释**：  时区。  **参数范围**：  不涉及。
	TimeZone *string `json:"time_zone,omitempty"`

	// **参数解释**：  付费模式。  **参数范围**：  不涉及。
	PayModel *string `json:"pay_model,omitempty"`

	// **参数解释**：  订单ID。  **参数范围**：  不涉及。
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**：  周期。  **参数范围**：  不涉及。
	Period *int32 `json:"period,omitempty"`

	// **参数解释**：  是否冻结。  **参数范围**：  不涉及。
	IsFrozen *bool `json:"is_frozen,omitempty"`

	// **参数解释**：  冻结时间。  **参数范围**：  不涉及。
	FrozenTime *string `json:"frozen_time,omitempty"`

	// **参数解释**：  锁状态。  **参数范围**：  不涉及。
	Actions *[]ActionInfo `json:"actions,omitempty"`

	// **参数解释**：  是否只有默认组。  **参数范围**：  不涉及。
	OnlyDefaultGroup *bool `json:"only_default_group,omitempty"`

	// **参数解释**：  组信息。  **参数范围**：  不涉及。
	Groups *[]DdmGroupInfo `json:"groups,omitempty"`

	// **参数解释**：  其他信息。  **参数范围**：  不涉及。
	ExtendMap map[string]string `json:"extend_map,omitempty"`

	// **参数解释**：  标签信息。  **参数范围**：  不涉及。
	TagsInfo *interface{} `json:"tags_info,omitempty"`

	// **参数解释**：  管理员账号。  **参数范围**：  不涉及。
	AdminUserName *string `json:"admin_user_name,omitempty"`

	// **参数解释**：  绑定eip信息。  **参数范围**：  不涉及。
	EipBindingInfo *interface{} `json:"eip_binding_info,omitempty"`

	// **参数解释**：  是否支持ssl。  **参数范围**：  不涉及。
	EnableSsl *int32 `json:"enable_ssl,omitempty"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}
