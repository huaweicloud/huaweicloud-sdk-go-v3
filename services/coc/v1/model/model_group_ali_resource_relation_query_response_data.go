package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type GroupAliResourceRelationQueryResponseData struct {

	// **参数解释：** 分组资源关联的id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** CloudCMDB分配的资源uuid。 **取值范围：** 不涉及。
	CmdbResourceId *string `json:"cmdb_resource_id,omitempty"`

	// **参数解释：** 分组id。 **取值范围：** 字符串，长度24个字符。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 分组名称。 **取值范围：** 字符串，长度3~60个字符。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释：** 在阿里云存的资源id。 **取值范围：** 不涉及。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释：** 资源名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 资源类型。 **取值范围：** 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 数据采集属性。 **取值范围：** 不涉及。
	IngestProperties *interface{} `json:"ingest_properties,omitempty"`

	// **参数解释：** 可用区id。 **取值范围：** 不涉及。
	ZoneId *string `json:"zone_id,omitempty"`

	// **参数解释：** 资源内网ip。 **取值范围：** 不涉及。
	InnerIp *string `json:"inner_ip,omitempty"`

	// **参数解释：** uniagent的id值。 **取值范围：** 不涉及。
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释：** uniagent的状态。 **取值范围：** - ONLINE：运行中。 - OFFLINE：异常。 - INSTALLING：安装中。 - FAILED：安装失败。 - UNINSTALLED：已卸载。 - null：未安装。
	AgentState *GroupAliResourceRelationQueryResponseDataAgentState `json:"agent_state,omitempty"`

	// **参数解释：** 区域id。 **取值范围：** 字符串，长度在0到64个字符之间。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 资源创建时间。 **取值范围：** 不涉及。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释：** 云厂商账户id。 **取值范围：** 不涉及。
	Datasource *string `json:"datasource,omitempty"`
}

func (o GroupAliResourceRelationQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupAliResourceRelationQueryResponseData struct{}"
	}

	return strings.Join([]string{"GroupAliResourceRelationQueryResponseData", string(data)}, " ")
}

type GroupAliResourceRelationQueryResponseDataAgentState struct {
	value string
}

type GroupAliResourceRelationQueryResponseDataAgentStateEnum struct {
	ONLINE      GroupAliResourceRelationQueryResponseDataAgentState
	OFFLINE     GroupAliResourceRelationQueryResponseDataAgentState
	INSTALLING  GroupAliResourceRelationQueryResponseDataAgentState
	FAILED      GroupAliResourceRelationQueryResponseDataAgentState
	UNINSTALLED GroupAliResourceRelationQueryResponseDataAgentState
}

func GetGroupAliResourceRelationQueryResponseDataAgentStateEnum() GroupAliResourceRelationQueryResponseDataAgentStateEnum {
	return GroupAliResourceRelationQueryResponseDataAgentStateEnum{
		ONLINE: GroupAliResourceRelationQueryResponseDataAgentState{
			value: "ONLINE",
		},
		OFFLINE: GroupAliResourceRelationQueryResponseDataAgentState{
			value: "OFFLINE",
		},
		INSTALLING: GroupAliResourceRelationQueryResponseDataAgentState{
			value: "INSTALLING",
		},
		FAILED: GroupAliResourceRelationQueryResponseDataAgentState{
			value: "FAILED",
		},
		UNINSTALLED: GroupAliResourceRelationQueryResponseDataAgentState{
			value: "UNINSTALLED",
		},
	}
}

func (c GroupAliResourceRelationQueryResponseDataAgentState) Value() string {
	return c.value
}

func (c GroupAliResourceRelationQueryResponseDataAgentState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GroupAliResourceRelationQueryResponseDataAgentState) UnmarshalJSON(b []byte) error {
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
