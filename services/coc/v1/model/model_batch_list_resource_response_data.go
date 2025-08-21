package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type BatchListResourceResponseData struct {

	// **参数解释：** CMDB分配的资源ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 云服务分配的资源ID。 **取值范围：** 字符串，长度在36个字符。
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释：** 租户id。 **取值范围：** 不涉及。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** 资源名称。 **取值范围：** 字符串，长度3到50个字符之间。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 云服务名称。 **取值范围：** 字符串，长度1到64个字符之间。
	Provider *string `json:"provider,omitempty"`

	// **参数解释：** 资源类型。 **取值范围：** 资源类型较多，根据实际业务选择资源类型、常用资源类型如下： - cloudservers：弹性云服务器。 - servers：裸金属服务器。 - clusters：云容器引擎。 - instances：云数据库。
	Type *string `json:"type,omitempty"`

	// **参数解释：** Openstack中的项目ID。 **取值范围：** 字符串，长度32个字符。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** region的子项目名称。 **取值范围：** 字符串，不超过255个字符。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释：** 区域id。 **取值范围：** 字符串，长度0到64个字符。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释：** 企业项目ID。 **取值范围：** 请选择[[企业管理](https://support.huaweicloud.com/usermanual-em/em_eps_qs_0400.html)](tag:hws)[[企业管理](https://support.huaweicloud.com/intl/zh-cn/usermanual-em/em_eps_qs_0400.html)](tag:hws_hk)中存在的项目ID。
	EpId *string `json:"ep_id,omitempty"`

	// **参数解释：** 企业项目名称。 **取值范围：** 由中文、英文字母、数字、下划线、中划线组成，且不能使用任何大小写形式的“default”，不超过255个字符。
	EpName *string `json:"ep_name,omitempty"`

	// **参数解释：** 资源标签。 **取值范围：** 不涉及。
	Tags *interface{} `json:"tags,omitempty"`

	// **参数解释：** uniagent的id值。 **取值范围：** 不涉及。
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释：** uniagent的状态。 **取值范围：** - ONLINE：运行中。 - OFFLINE：异常。 - INSTALLING：安装中。 - FAILED：安装失败。 - UNINSTALLED：已卸载。 - null：未安装。
	AgentState *BatchListResourceResponseDataAgentState `json:"agent_state,omitempty"`

	// **参数解释：** 资源详细属性。 **取值范围：** 不涉及。
	Properties *interface{} `json:"properties,omitempty"`

	// **参数解释：** 采集属性。 **取值范围：** 不涉及。
	IngestProperties map[string]string `json:"ingest_properties,omitempty"`

	// **参数解释：** 是否已托管。 **取值范围：** - true：已经托管。 - false：未托管。
	IsDelegated *bool `json:"is_delegated,omitempty"`

	// **参数解释：** 资源内网ip。 **取值范围：** 不涉及。
	InnerIp *string `json:"inner_ip,omitempty"`

	// **参数解释：** 用户定义资源是否可运维实例。 **取值范围：** - ENABLE：可运维实例。 - DISABLE：不可运维实例operable字段不存在。
	Operable *string `json:"operable,omitempty"`

	// **参数解释：** 是否已被指定分组关联。 **取值范围：** - true：已被指定分组关联。 - false：未被指定分组关联。
	IsAssociateGroup *bool `json:"is_associate_group,omitempty"`

	// **参数解释：** 资源所关联的分组信息组成的列表。 **取值范围：** 不涉及。
	AssociatedGroupList *[]string `json:"associated_group_list,omitempty"`

	// **参数解释：** 创建时间，参考ISO8601标准格式。 **取值范围：** 不涉及。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// **参数解释：** 修改时间，参考ISO8601标准格式。 **取值范围：** 不涉及。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`
}

func (o BatchListResourceResponseData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListResourceResponseData struct{}"
	}

	return strings.Join([]string{"BatchListResourceResponseData", string(data)}, " ")
}

type BatchListResourceResponseDataAgentState struct {
	value string
}

type BatchListResourceResponseDataAgentStateEnum struct {
	ONLINE      BatchListResourceResponseDataAgentState
	OFFLINE     BatchListResourceResponseDataAgentState
	INSTALLING  BatchListResourceResponseDataAgentState
	FAILED      BatchListResourceResponseDataAgentState
	UNINSTALLED BatchListResourceResponseDataAgentState
}

func GetBatchListResourceResponseDataAgentStateEnum() BatchListResourceResponseDataAgentStateEnum {
	return BatchListResourceResponseDataAgentStateEnum{
		ONLINE: BatchListResourceResponseDataAgentState{
			value: "ONLINE",
		},
		OFFLINE: BatchListResourceResponseDataAgentState{
			value: "OFFLINE",
		},
		INSTALLING: BatchListResourceResponseDataAgentState{
			value: "INSTALLING",
		},
		FAILED: BatchListResourceResponseDataAgentState{
			value: "FAILED",
		},
		UNINSTALLED: BatchListResourceResponseDataAgentState{
			value: "UNINSTALLED",
		},
	}
}

func (c BatchListResourceResponseDataAgentState) Value() string {
	return c.value
}

func (c BatchListResourceResponseDataAgentState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchListResourceResponseDataAgentState) UnmarshalJSON(b []byte) error {
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
