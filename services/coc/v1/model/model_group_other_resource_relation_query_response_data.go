package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type GroupOtherResourceRelationQueryResponseData struct {

	// **参数解释：** CMDB生成uuid。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 租户id。 **取值范围：** 字符串，长度32个字符。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** CloudCMDB分配的资源uuid。 **取值范围：** 不涉及。
	CmdbResourceId *string `json:"cmdb_resource_id,omitempty"`

	// **参数解释：** 分组id。 **取值范围：** 字符串，长度24个字符。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 分组名称。 **取值范围：** 字符串，长度3~60个字符。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释：** 资源id。 **取值范围：** 字符串，长度36个字符。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释：** 资源名称。 **取值范围：** 字符串，长度3到50个字符串。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 资源类型。 **取值范围：** - PM：物理机。 - VM：虚拟机。
	Type *GroupOtherResourceRelationQueryResponseDataType `json:"type,omitempty"`

	// **参数解释：** 是否已托管。 **取值范围：** - true：已经托管。 - false：未托管。
	IsDelegated *bool `json:"is_delegated,omitempty"`

	// **参数解释：** 区域id。 **取值范围：** 字符串，长度在0到64个字符之间。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 资源内网ip。 **取值范围：** 不涉及。
	InnerIp *string `json:"inner_ip,omitempty"`

	// **参数解释：** 可用区id。 **取值范围：** 不涉及。
	ZoneId *string `json:"zone_id,omitempty"`

	// **参数解释：** 是否为主机。 **取值范围：** - true：为主机。 - false： 非主机。
	IsHost *bool `json:"is_host,omitempty"`

	// **参数解释：** uniagent的id值。 **取值范围：** 不涉及。
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释：** uniagent的状态。 **取值范围：** - ONLINE：运行中。 - OFFLINE：异常。 - INSTALLING：安装中。 - FAILED：安装失败。 - UNINSTALLED：已卸载。 - null：未安装。
	AgentState *GroupOtherResourceRelationQueryResponseDataAgentState `json:"agent_state,omitempty"`

	// **参数解释：** 云厂商账户id。 **取值范围：** 字符串，长度在0到24个字符之间。
	Datasource *string `json:"datasource,omitempty"`

	// **参数解释：** 描述。 **取值范围：** 字符串，长度在1到256个字符之间。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 属性信息。 **取值范围：** 不涉及。
	Properties map[string]interface{} `json:"properties,omitempty"`

	// **参数解释：** 数据采集属性，描述视图系统采集该资源时所记录的附加信息。 **取值范围：** 不涉及。
	IngestProperties map[string]interface{} `json:"ingest_properties,omitempty"`

	// **参数解释：** 设备标识。 **取值范围：** 字符串，长度3到50个字符之间。
	Xrn *string `json:"xrn,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 不涉及。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o GroupOtherResourceRelationQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupOtherResourceRelationQueryResponseData struct{}"
	}

	return strings.Join([]string{"GroupOtherResourceRelationQueryResponseData", string(data)}, " ")
}

type GroupOtherResourceRelationQueryResponseDataType struct {
	value string
}

type GroupOtherResourceRelationQueryResponseDataTypeEnum struct {
	PM GroupOtherResourceRelationQueryResponseDataType
	VM GroupOtherResourceRelationQueryResponseDataType
}

func GetGroupOtherResourceRelationQueryResponseDataTypeEnum() GroupOtherResourceRelationQueryResponseDataTypeEnum {
	return GroupOtherResourceRelationQueryResponseDataTypeEnum{
		PM: GroupOtherResourceRelationQueryResponseDataType{
			value: "PM",
		},
		VM: GroupOtherResourceRelationQueryResponseDataType{
			value: "VM",
		},
	}
}

func (c GroupOtherResourceRelationQueryResponseDataType) Value() string {
	return c.value
}

func (c GroupOtherResourceRelationQueryResponseDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GroupOtherResourceRelationQueryResponseDataType) UnmarshalJSON(b []byte) error {
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

type GroupOtherResourceRelationQueryResponseDataAgentState struct {
	value string
}

type GroupOtherResourceRelationQueryResponseDataAgentStateEnum struct {
	ONLINE      GroupOtherResourceRelationQueryResponseDataAgentState
	OFFLINE     GroupOtherResourceRelationQueryResponseDataAgentState
	INSTALLING  GroupOtherResourceRelationQueryResponseDataAgentState
	FAILED      GroupOtherResourceRelationQueryResponseDataAgentState
	UNINSTALLED GroupOtherResourceRelationQueryResponseDataAgentState
}

func GetGroupOtherResourceRelationQueryResponseDataAgentStateEnum() GroupOtherResourceRelationQueryResponseDataAgentStateEnum {
	return GroupOtherResourceRelationQueryResponseDataAgentStateEnum{
		ONLINE: GroupOtherResourceRelationQueryResponseDataAgentState{
			value: "ONLINE",
		},
		OFFLINE: GroupOtherResourceRelationQueryResponseDataAgentState{
			value: "OFFLINE",
		},
		INSTALLING: GroupOtherResourceRelationQueryResponseDataAgentState{
			value: "INSTALLING",
		},
		FAILED: GroupOtherResourceRelationQueryResponseDataAgentState{
			value: "FAILED",
		},
		UNINSTALLED: GroupOtherResourceRelationQueryResponseDataAgentState{
			value: "UNINSTALLED",
		},
	}
}

func (c GroupOtherResourceRelationQueryResponseDataAgentState) Value() string {
	return c.value
}

func (c GroupOtherResourceRelationQueryResponseDataAgentState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GroupOtherResourceRelationQueryResponseDataAgentState) UnmarshalJSON(b []byte) error {
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
