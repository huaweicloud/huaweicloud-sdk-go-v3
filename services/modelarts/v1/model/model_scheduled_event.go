package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ScheduledEvent struct {

	// **参数解释**：计划事件ID，取值查询计划事件列表接口的event_id字段。 系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，长度小于63。 **取值范围**：不涉及。
	Id string `json:"id"`

	// **参数解释**：事件分类。 **取值范围**：不涉及。
	Catalog *ScheduledEventCatalog `json:"catalog,omitempty"`

	// **参数解释**：超节点标识。 **约束限制**：不涉及。 **取值范围**：true / false **默认取值**：false
	IsHyperNode *string `json:"isHyperNode,omitempty"`

	// **参数解释**：事件类型。 **取值范围**：可选值如下： - system-maintenance：系统维护 - localdisk-recovery：本地盘恢复 - node_reboot：节点重启 - operation-request：运维授权 - node_maintenance：超节点维护 - node_redeploy：超节点重部署 - node_localdisk_recovery 超节点本地盘恢复
	Type *ScheduledEventType `json:"type,omitempty"`

	// **参数解释**：对计划事件的描述信息。系统自动生成。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：事件状态。 **取值范围**：可选择值如下： - inquiring: 待授权 - scheduled: 待执行 - executing: 执行中 - completed: 执行成功 - failed: 执行失败 - canceled: 取消 - part-completed: 部分成功
	State ScheduledEventState `json:"state"`

	// **参数解释**：事件发布时间。 **取值范围**：不涉及。
	PublishTime string `json:"publishTime"`

	// **参数解释**：事件开始时间。 **取值范围**：不涉及。
	StartTime *string `json:"startTime,omitempty"`

	// **参数解释**：事件完成时间。 **取值范围**：不涉及。
	FinishTime *string `json:"finishTime,omitempty"`

	// **参数解释**：事件计划执行开始时间，格式为UTC时间字符串：2025-09-15T07:02:30Z。 **取值范围**：不涉及。
	NotBefore *string `json:"notBefore,omitempty"`

	// **参数解释**：提示信息。 **取值范围**：不涉及。
	ProbeMsg *string `json:"probeMsg,omitempty"`

	ExecuteOptions *ExecuteOptions `json:"executeOptions,omitempty"`

	// **参数解释**：节点重部署类型列表，标识该计划事件支持的重部署方式。 **约束限制**：不涉及。 **取值范围**：可选值如下： - SOFT：表示支持重部署 - HARD：表示支持强制重部署 **默认取值**：不涉及。
	RedeployType []string `json:"redeployType"`

	// **参数解释**：事件系统自动授权执行时间，格式为UTC时间字符串：2025-09-15T07:02:30Z。 **取值范围**：不涉及。
	AutoRecoveryCheckTime *string `json:"autoRecoveryCheckTime,omitempty"`

	// **参数解释**：事件关联节点信息列表。 **取值范围**：不涉及。
	Nodes *[]NodeEvent `json:"nodes,omitempty"`

	// **参数解释**：节点ID。计算服务系统自动生成的实例ID，长度小于63。 **取值范围**：不涉及。 **约束限制**：待废弃: 请使用 nodes[0].InstanceId。
	InstanceId string `json:"instanceId"`

	// **参数解释**：节点名称，取值自节点详情的metadata.name字段。系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，小于63个字符。 **取值范围**：不涉及。 **约束限制**：待废弃: 请使用 nodes[0].nodeName。
	NodeName string `json:"nodeName"`

	// **参数解释**：节点类型归属。 **取值范围**可选择值如下： - devserver：lite-server节点  - lite-cluster lite池  - standard 标准池 **约束限制**：待废弃: 请使用 nodes[0].instanceType。
	InstanceType ScheduledEventInstanceType `json:"instanceType"`

	// **参数解释**：超节点场景对应ECS服务器ID。计算服务系统自动生成的实例ID，长度小于63。  **取值范围**：超节点涉及, snt9b 不涉及。 **约束限制**：待废弃: 请使用 nodes[0].underLyingInstanceId。
	UnderLyingInstanceId *string `json:"underLyingInstanceId,omitempty"`

	// **参数解释**：规格是否支持重部署。 **取值范围**：true/false。 **约束限制**：待废弃: 请使用 nodes[0].flavorSupportRedeploy。
	FlavorSupportRedeploy *bool `json:"flavorSupportRedeploy,omitempty"`

	// **参数解释**：资源池名称, lite-cluster、standard才具有，取值自资源池详情的metadata.name字段。系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，小于63个字符。 **取值范围**：不涉及。 **约束限制**：待废弃: 请使用 nodes[0].poolName。
	PoolName *string `json:"poolName,omitempty"`

	// **参数解释**：资源池对外显示的名称, lite-cluster、standard才具有，取值自资源池详情的metadata.labels字段。只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为[36-63]个字符。 **取值范围**：不涉及。 **约束限制**：待废弃: 请使用 nodes[0].poolDisplayName。
	PoolDisplayName *string `json:"poolDisplayName,omitempty"`

	// **参数解释**：重部署执行的实时进度信息，系统自动生成。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。 **约束限制**：待废弃: 请使用 nodes[0].jobProgress。
	JobProgress *string `json:"jobProgress,omitempty"`
}

func (o ScheduledEvent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduledEvent struct{}"
	}

	return strings.Join([]string{"ScheduledEvent", string(data)}, " ")
}

type ScheduledEventCatalog struct {
	value string
}

type ScheduledEventCatalogEnum struct {
	SOFTWARE ScheduledEventCatalog
	HARDWARE ScheduledEventCatalog
}

func GetScheduledEventCatalogEnum() ScheduledEventCatalogEnum {
	return ScheduledEventCatalogEnum{
		SOFTWARE: ScheduledEventCatalog{
			value: "software",
		},
		HARDWARE: ScheduledEventCatalog{
			value: "hardware",
		},
	}
}

func (c ScheduledEventCatalog) Value() string {
	return c.value
}

func (c ScheduledEventCatalog) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledEventCatalog) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ScheduledEventType struct {
	value string
}

type ScheduledEventTypeEnum struct {
	SYSTEM_MAINTENANCE      ScheduledEventType
	LOCALDISK_RECOVERY      ScheduledEventType
	NODE_REBOOT             ScheduledEventType
	OPERATION_REQUEST       ScheduledEventType
	NODE_MAINTENANCE        ScheduledEventType
	NODE_REDEPLOY           ScheduledEventType
	NODE_LOCALDISK_RECOVERY ScheduledEventType
}

func GetScheduledEventTypeEnum() ScheduledEventTypeEnum {
	return ScheduledEventTypeEnum{
		SYSTEM_MAINTENANCE: ScheduledEventType{
			value: "system-maintenance",
		},
		LOCALDISK_RECOVERY: ScheduledEventType{
			value: "localdisk-recovery",
		},
		NODE_REBOOT: ScheduledEventType{
			value: "node_reboot",
		},
		OPERATION_REQUEST: ScheduledEventType{
			value: "operation-request",
		},
		NODE_MAINTENANCE: ScheduledEventType{
			value: "node_maintenance",
		},
		NODE_REDEPLOY: ScheduledEventType{
			value: "node_redeploy",
		},
		NODE_LOCALDISK_RECOVERY: ScheduledEventType{
			value: "node_localdisk_recovery",
		},
	}
}

func (c ScheduledEventType) Value() string {
	return c.value
}

func (c ScheduledEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledEventType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ScheduledEventState struct {
	value string
}

type ScheduledEventStateEnum struct {
	INQUIRING      ScheduledEventState
	SCHEDULED      ScheduledEventState
	EXECUTING      ScheduledEventState
	COMPLETED      ScheduledEventState
	FAILED         ScheduledEventState
	CANCELED       ScheduledEventState
	PART_COMPLETED ScheduledEventState
}

func GetScheduledEventStateEnum() ScheduledEventStateEnum {
	return ScheduledEventStateEnum{
		INQUIRING: ScheduledEventState{
			value: "inquiring",
		},
		SCHEDULED: ScheduledEventState{
			value: "scheduled",
		},
		EXECUTING: ScheduledEventState{
			value: "executing",
		},
		COMPLETED: ScheduledEventState{
			value: "completed",
		},
		FAILED: ScheduledEventState{
			value: "failed",
		},
		CANCELED: ScheduledEventState{
			value: "canceled",
		},
		PART_COMPLETED: ScheduledEventState{
			value: "part-completed",
		},
	}
}

func (c ScheduledEventState) Value() string {
	return c.value
}

func (c ScheduledEventState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledEventState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ScheduledEventInstanceType struct {
	value string
}

type ScheduledEventInstanceTypeEnum struct {
	DEVSERVER    ScheduledEventInstanceType
	LITE_CLUSTER ScheduledEventInstanceType
	STANDARD     ScheduledEventInstanceType
}

func GetScheduledEventInstanceTypeEnum() ScheduledEventInstanceTypeEnum {
	return ScheduledEventInstanceTypeEnum{
		DEVSERVER: ScheduledEventInstanceType{
			value: "devserver",
		},
		LITE_CLUSTER: ScheduledEventInstanceType{
			value: "lite-cluster",
		},
		STANDARD: ScheduledEventInstanceType{
			value: "standard",
		},
	}
}

func (c ScheduledEventInstanceType) Value() string {
	return c.value
}

func (c ScheduledEventInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduledEventInstanceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
