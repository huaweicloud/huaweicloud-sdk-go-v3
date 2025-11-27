package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type GroupRmsResourceRelationQueryResponseData struct {

	// **参数解释：** 分组资源关联的id。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** CloudCMDB分配的资源uuid。
	CmdbResourceId *string `json:"cmdb_resource_id,omitempty"`

	// **参数解释：** 分组id。 **取值范围：** 不涉及。
	GroupId *string `json:"group_id,omitempty"`

	// **参数解释：** 分组名称。 **取值范围：** 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释：** 应用id。 **取值范围：** 长度24。
	ApplicationId *string `json:"application_id,omitempty"`

	// **参数解释：** 资源id。 **取值范围：** 长度36
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释：** 资源名称。 **取值范围：** 长度3到50之间。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 企业项目id。 **取值范围：** 不涉及。
	EpId *string `json:"ep_id,omitempty"`

	// **参数解释：** 项目id。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 企业项目名称。 **取值范围：** 不涉及。
	EpName *string `json:"ep_name,omitempty"`

	// **参数解释：** 云服务名称。 **取值范围：** 不涉及。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 租户id。 **取值范围：** 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 资源类型。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 区域id。 **取值范围：** 字符串，长度在0到64个字符之间。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 资源内网ip。 **取值范围：** 不涉及。
	InnerIp *string `json:"inner_ip,omitempty"`

	// **参数解释：** uniagent的id值。 **取值范围：** 不涉及。
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释：** uniagent的状态。 **取值范围：** - ONLINE：运行中。 - OFFLINE：异常。 - INSTALLING：安装中。 - FAILED：安装失败。 - UNINSTALLED：已卸载。 - null：未安装。
	AgentState *GroupRmsResourceRelationQueryResponseDataAgentState `json:"agent_state,omitempty"`

	// **参数解释：** 标签键值对。 **取值范围：** 不涉及。
	Tags *[]GroupRmsResourceRelationQueryResponseTags `json:"tags,omitempty"`

	// **参数解释：** 数据采集属性。 **取值范围：** 不涉及。
	IngestProperties *interface{} `json:"ingest_properties,omitempty"`

	// **参数解释：** 属性。 **取值范围：** 不涉及。
	Properties map[string]string `json:"properties,omitempty"`

	// **参数解释：** 用户定义资源是否可运维实例。 **取值范围：** - ENABLE：可运维。 - DISABLE：不可运维实例operable字段不存在。
	Operable *string `json:"operable,omitempty"`

	// **参数解释：** 创建时间，参考ISO8601标准格式。 **取值范围：** 不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`
}

func (o GroupRmsResourceRelationQueryResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupRmsResourceRelationQueryResponseData struct{}"
	}

	return strings.Join([]string{"GroupRmsResourceRelationQueryResponseData", string(data)}, " ")
}

type GroupRmsResourceRelationQueryResponseDataAgentState struct {
	value string
}

type GroupRmsResourceRelationQueryResponseDataAgentStateEnum struct {
	ONLINE      GroupRmsResourceRelationQueryResponseDataAgentState
	OFFLINE     GroupRmsResourceRelationQueryResponseDataAgentState
	INSTALLING  GroupRmsResourceRelationQueryResponseDataAgentState
	FAILED      GroupRmsResourceRelationQueryResponseDataAgentState
	UNINSTALLED GroupRmsResourceRelationQueryResponseDataAgentState
}

func GetGroupRmsResourceRelationQueryResponseDataAgentStateEnum() GroupRmsResourceRelationQueryResponseDataAgentStateEnum {
	return GroupRmsResourceRelationQueryResponseDataAgentStateEnum{
		ONLINE: GroupRmsResourceRelationQueryResponseDataAgentState{
			value: "ONLINE",
		},
		OFFLINE: GroupRmsResourceRelationQueryResponseDataAgentState{
			value: "OFFLINE",
		},
		INSTALLING: GroupRmsResourceRelationQueryResponseDataAgentState{
			value: "INSTALLING",
		},
		FAILED: GroupRmsResourceRelationQueryResponseDataAgentState{
			value: "FAILED",
		},
		UNINSTALLED: GroupRmsResourceRelationQueryResponseDataAgentState{
			value: "UNINSTALLED",
		},
	}
}

func (c GroupRmsResourceRelationQueryResponseDataAgentState) Value() string {
	return c.value
}

func (c GroupRmsResourceRelationQueryResponseDataAgentState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GroupRmsResourceRelationQueryResponseDataAgentState) UnmarshalJSON(b []byte) error {
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
