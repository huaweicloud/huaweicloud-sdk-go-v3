package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AcceptScheduledEventResponse Response Object
type AcceptScheduledEventResponse struct {

	// **参数解释**：计划事件ID，取值查询计划事件列表接口的event_id字段。 系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，长度小于63。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：事件分类。 **取值范围**：可选值如下： - hardware：硬件维修。 - software：软件维修。
	Catalog *AcceptScheduledEventResponseCatalog `json:"catalog,omitempty"`

	// **参数解释**：事件类型。 **取值范围**：可选值如下： - system-maintenance：系统维护 - localdisk-recovery：本地盘恢复 - node_reboot：节点重启 - operation-request：运维授权 - node_maintenance：超节点维护 - node_redeploy：超节点重部署 - node_localdisk_recovery 超节点本地盘恢复。
	Type *AcceptScheduledEventResponseType `json:"type,omitempty"`

	// **参数解释**：对计划事件的描述信息。系统自动生成。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：事件状态。 **取值范围**：可选择值如下： - inquiring: 待授权, - scheduled: 待执行, - executing: 执行中, - completed: 执行成功 - failed: 执行失败 - canceled: 取消
	State *AcceptScheduledEventResponseState `json:"state,omitempty"`

	// **参数解释**：节点类型归属。 **取值范围**可选择值如下： - devserver：lite-server节点 - lite-cluster lite池 - standard 标准池
	InstanceType *AcceptScheduledEventResponseInstanceType `json:"instanceType,omitempty"`

	// **参数解释**：服务器ID。计算服务系统自动生成的实例ID，长度小于63。 **取值范围**：不涉及。
	InstanceId *string `json:"instanceId,omitempty"`

	// **参数解释**：节点名称，取值自节点详情的metadata.name字段。系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，小于63个字符。 **取值范围**：不涉及。
	NodeName *string `json:"nodeName,omitempty"`

	// **参数解释**：资源池名称, lite-cluster、standard才具有，取值自资源池详情的metadata.name字段。系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，小于63个字符。 **取值范围**：不涉及。
	PoolName *string `json:"poolName,omitempty"`

	// **参数解释**：资源池对外显示的名称, lite-cluster、standard才具有，取值自资源池详情的metadata.name字段。只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为[36-63]个字符。 **取值范围**：不涉及。
	PoolDisplayName *string `json:"poolDisplayName,omitempty"`

	// **参数解释**：事件发布时间。 **约束限制**：格式为UTC时间字符串：2025-09-15T07:02:30Z。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PublishTime *string `json:"publishTime,omitempty"`

	// **参数解释**：事件开始时间。 **约束限制**：格式为UTC时间字符串：2025-09-15T07:02:30Z。 **取值范围**：不涉及。 **默认取值**：不涉及。
	StartTime *string `json:"startTime,omitempty"`

	// **参数解释**：事件完成时间。 **约束限制**：格式为UTC时间字符串：2025-09-15T07:02:30Z。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FinishTime *string `json:"finishTime,omitempty"`

	// **参数解释**：事件计划执行开始时间，格式为UTC时间字符串：2025-09-15T07:02:30Z。 **约束限制**：大于当前时间。 **取值范围**：不涉及。 **默认取值**：不填表示立即执行。
	NotBefore *string `json:"notBefore,omitempty"`

	// **参数解释**：提示信息。 **约束限制**：系统自动生成，只能以小写字母开头，数字、中划线组成，不能以中划线结尾，长度小于63字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ProbeMsg *string `json:"probeMsg,omitempty"`

	// **参数解释**：节点的重部署类型。 **约束限制**：不涉及。 **取值范围**：可选值如下：- HARD：表示支持强制重部署, - SOFT：表示支持重部署 **默认取值**：不涉及。
	RedeployType *[]string `json:"redeployType,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AcceptScheduledEventResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AcceptScheduledEventResponse struct{}"
	}

	return strings.Join([]string{"AcceptScheduledEventResponse", string(data)}, " ")
}

type AcceptScheduledEventResponseCatalog struct {
	value string
}

type AcceptScheduledEventResponseCatalogEnum struct {
	SOFTWARE AcceptScheduledEventResponseCatalog
	HARDWARE AcceptScheduledEventResponseCatalog
}

func GetAcceptScheduledEventResponseCatalogEnum() AcceptScheduledEventResponseCatalogEnum {
	return AcceptScheduledEventResponseCatalogEnum{
		SOFTWARE: AcceptScheduledEventResponseCatalog{
			value: "software",
		},
		HARDWARE: AcceptScheduledEventResponseCatalog{
			value: "hardware",
		},
	}
}

func (c AcceptScheduledEventResponseCatalog) Value() string {
	return c.value
}

func (c AcceptScheduledEventResponseCatalog) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptScheduledEventResponseCatalog) UnmarshalJSON(b []byte) error {
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

type AcceptScheduledEventResponseType struct {
	value string
}

type AcceptScheduledEventResponseTypeEnum struct {
	SYSTEM_MAINTENANCE      AcceptScheduledEventResponseType
	LOCALDISK_RECOVERY      AcceptScheduledEventResponseType
	NODE_REBOOT             AcceptScheduledEventResponseType
	OPERATION_REQUEST       AcceptScheduledEventResponseType
	NODE_MAINTENANCE        AcceptScheduledEventResponseType
	NODE_REDEPLOY           AcceptScheduledEventResponseType
	NODE_LOCALDISK_RECOVERY AcceptScheduledEventResponseType
}

func GetAcceptScheduledEventResponseTypeEnum() AcceptScheduledEventResponseTypeEnum {
	return AcceptScheduledEventResponseTypeEnum{
		SYSTEM_MAINTENANCE: AcceptScheduledEventResponseType{
			value: "system-maintenance",
		},
		LOCALDISK_RECOVERY: AcceptScheduledEventResponseType{
			value: "localdisk-recovery",
		},
		NODE_REBOOT: AcceptScheduledEventResponseType{
			value: "node_reboot",
		},
		OPERATION_REQUEST: AcceptScheduledEventResponseType{
			value: "operation-request",
		},
		NODE_MAINTENANCE: AcceptScheduledEventResponseType{
			value: "node_maintenance",
		},
		NODE_REDEPLOY: AcceptScheduledEventResponseType{
			value: "node_redeploy",
		},
		NODE_LOCALDISK_RECOVERY: AcceptScheduledEventResponseType{
			value: "node_localdisk_recovery",
		},
	}
}

func (c AcceptScheduledEventResponseType) Value() string {
	return c.value
}

func (c AcceptScheduledEventResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptScheduledEventResponseType) UnmarshalJSON(b []byte) error {
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

type AcceptScheduledEventResponseState struct {
	value string
}

type AcceptScheduledEventResponseStateEnum struct {
	INQUIRING AcceptScheduledEventResponseState
	SCHEDULED AcceptScheduledEventResponseState
	EXECUTING AcceptScheduledEventResponseState
	COMPLETED AcceptScheduledEventResponseState
	FAILED    AcceptScheduledEventResponseState
	CANCELED  AcceptScheduledEventResponseState
}

func GetAcceptScheduledEventResponseStateEnum() AcceptScheduledEventResponseStateEnum {
	return AcceptScheduledEventResponseStateEnum{
		INQUIRING: AcceptScheduledEventResponseState{
			value: "inquiring",
		},
		SCHEDULED: AcceptScheduledEventResponseState{
			value: "scheduled",
		},
		EXECUTING: AcceptScheduledEventResponseState{
			value: "executing",
		},
		COMPLETED: AcceptScheduledEventResponseState{
			value: "completed",
		},
		FAILED: AcceptScheduledEventResponseState{
			value: "failed",
		},
		CANCELED: AcceptScheduledEventResponseState{
			value: "canceled",
		},
	}
}

func (c AcceptScheduledEventResponseState) Value() string {
	return c.value
}

func (c AcceptScheduledEventResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptScheduledEventResponseState) UnmarshalJSON(b []byte) error {
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

type AcceptScheduledEventResponseInstanceType struct {
	value string
}

type AcceptScheduledEventResponseInstanceTypeEnum struct {
	DEVSERVER    AcceptScheduledEventResponseInstanceType
	LITE_CLUSTER AcceptScheduledEventResponseInstanceType
	STANDARD     AcceptScheduledEventResponseInstanceType
}

func GetAcceptScheduledEventResponseInstanceTypeEnum() AcceptScheduledEventResponseInstanceTypeEnum {
	return AcceptScheduledEventResponseInstanceTypeEnum{
		DEVSERVER: AcceptScheduledEventResponseInstanceType{
			value: "devserver",
		},
		LITE_CLUSTER: AcceptScheduledEventResponseInstanceType{
			value: "lite-cluster",
		},
		STANDARD: AcceptScheduledEventResponseInstanceType{
			value: "standard",
		},
	}
}

func (c AcceptScheduledEventResponseInstanceType) Value() string {
	return c.value
}

func (c AcceptScheduledEventResponseInstanceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AcceptScheduledEventResponseInstanceType) UnmarshalJSON(b []byte) error {
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
