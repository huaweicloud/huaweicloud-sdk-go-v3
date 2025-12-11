package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ScheduleTask struct {

	// 任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例名称。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例状态。 取值： - 值为“createfail”，表示实例创建失败。 - 值为“creating”，表示实例创建中。 - 值为“normal”，表示实例正常。 - 值为“abnormal”，表示实例异常。 - 值为“deleted”，表示实例已删除。
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 租户在某一region下的project ID。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 任务名称。  **取值范围**：   - PROXY_VERSION_UPGRADE：表示数据库代理版本升级。   - VERSION_UPGRADE：表示实例版本升级。   - RESIZE_FLAVOR：表示实例规格变更。   - REBOOT_NODE：表示重启节点。   - REBOOT_INSTANCE：表示重启实例。
	JobName *string `json:"job_name,omitempty"`

	// 任务创建时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。 说明：创建时返回值为空，数据库实例创建成功后该值不为空
	CreateTime *string `json:"create_time,omitempty"`

	// 任务开始时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。 说明：创建时返回值为空，数据库实例创建成功后该值不为空
	StartTime *string `json:"start_time,omitempty"`

	// 任务结束时间，格式为\"yyyy-mm-ddThh:mm:ssZ\"。 其中，T指某个时间的开始；Z指时区偏移量，例如偏移1个小时显示为+0100。 说明：创建时返回值为空，数据库实例创建成功后该值不为空
	EndTime *string `json:"end_time,omitempty"`

	// 任务执行状态。 取值： - 值为“Pending”，表示延时任务，未执行。 - 值为“Running”，表示任务正在执行。 - 值为“Completed”，表示任务执行成功。 - 值为“Failed”，表示任务执行失败。
	JobStatus *string `json:"job_status,omitempty"`

	// 数据库类型。
	DatastoreType *string `json:"datastore_type,omitempty"`

	// 实例配置相关信息，比如规格等。
	TargetConfig *interface{} `json:"target_config,omitempty"`

	// **参数解释**：  数据库代理ID，严格匹配UUID规则。  **取值范围**：  只能由英文字母、数字组成，后缀为po01，长度为36个字符。
	ProxyId *string `json:"proxy_id,omitempty"`

	// **参数解释**：  数据库代理名称。  **取值范围**：  不涉及。
	ProxyName *string `json:"proxy_name,omitempty"`
}

func (o ScheduleTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleTask struct{}"
	}

	return strings.Join([]string{"ScheduleTask", string(data)}, " ")
}
