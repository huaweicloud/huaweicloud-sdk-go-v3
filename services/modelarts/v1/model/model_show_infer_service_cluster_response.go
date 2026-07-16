package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInferServiceClusterResponse Response Object
type ShowInferServiceClusterResponse struct {

	// **参数解释：** 工作空间ID。未创建工作空间时默认值为“0”，存在创建并使用的工作空间，以实际取值为准。 **取值范围：** 不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释：** 资源池cceID。 **取值范围：** 不涉及。
	LogicClusterId *string `json:"logic_cluster_id,omitempty"`

	// **参数解释：** 资源当前状态。 **取值范围：** - ACTIVE ：开启。 - PENDING ：待处理。 - INITIALIZING ：初始化中。 - INITIALIZE_FAILED ：初始化失败。 - DELETED ：已删除。 - DELETING ：删除中。 - DELETE_FAILED ：删除失败。 - MIGRATING : 迁移中。
	Status *ShowInferServiceClusterResponseStatus `json:"status,omitempty"`

	// **参数解释：** 专属资源池ID。 **取值范围：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 专属池类型。 **取值范围：** - MANAGED ：公共。 - MANAGED_ROMA ：公共。 - DEDICATED ：专属。 - DEDICATED_ROMA ：专属。
	Type *ShowInferServiceClusterResponseType `json:"type,omitempty"`

	// **参数解释：** 资源池类型。
	ResourceCategories *[]string `json:"resource_categories,omitempty"`

	// **参数解释：** [用户项目ID](tag:hws,hws_hk,fcs,fcs_super)[资源空间ID](tag:hcs,hcs_sm)。获取方法请参见[[获取项目ID和名称](modelarts_03_0147.xml)](tag:hws,hws_hk,fcs,fcs_super)[[获取资源空间ID和名称](modelarts_03_0147.xml)](tag:hcs,hcs_sm)。 **取值范围：** 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释：** 资源池启用的时间，UTC毫秒。 **取值范围：** 不涉及。
	CreateAt *int64 `json:"create_at,omitempty"`

	// **参数解释：** 资源池最后更新的时间，UTC毫秒。 **取值范围：** 不涉及。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释：** 当前专属池支持的规格。
	Flavors        *[]InferFlavor `json:"flavors,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowInferServiceClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInferServiceClusterResponse struct{}"
	}

	return strings.Join([]string{"ShowInferServiceClusterResponse", string(data)}, " ")
}

type ShowInferServiceClusterResponseStatus struct {
	value string
}

type ShowInferServiceClusterResponseStatusEnum struct {
	ACTIVE            ShowInferServiceClusterResponseStatus
	PENDING           ShowInferServiceClusterResponseStatus
	INITIALIZING      ShowInferServiceClusterResponseStatus
	INITIALIZE_FAILED ShowInferServiceClusterResponseStatus
	DELETED           ShowInferServiceClusterResponseStatus
	DELETING          ShowInferServiceClusterResponseStatus
	DELETE_FAILED     ShowInferServiceClusterResponseStatus
	MIGRATING         ShowInferServiceClusterResponseStatus
}

func GetShowInferServiceClusterResponseStatusEnum() ShowInferServiceClusterResponseStatusEnum {
	return ShowInferServiceClusterResponseStatusEnum{
		ACTIVE: ShowInferServiceClusterResponseStatus{
			value: "ACTIVE",
		},
		PENDING: ShowInferServiceClusterResponseStatus{
			value: "PENDING",
		},
		INITIALIZING: ShowInferServiceClusterResponseStatus{
			value: "INITIALIZING",
		},
		INITIALIZE_FAILED: ShowInferServiceClusterResponseStatus{
			value: "INITIALIZE_FAILED",
		},
		DELETED: ShowInferServiceClusterResponseStatus{
			value: "DELETED",
		},
		DELETING: ShowInferServiceClusterResponseStatus{
			value: "DELETING",
		},
		DELETE_FAILED: ShowInferServiceClusterResponseStatus{
			value: "DELETE_FAILED",
		},
		MIGRATING: ShowInferServiceClusterResponseStatus{
			value: "MIGRATING",
		},
	}
}

func (c ShowInferServiceClusterResponseStatus) Value() string {
	return c.value
}

func (c ShowInferServiceClusterResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInferServiceClusterResponseStatus) UnmarshalJSON(b []byte) error {
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

type ShowInferServiceClusterResponseType struct {
	value string
}

type ShowInferServiceClusterResponseTypeEnum struct {
	MANAGED        ShowInferServiceClusterResponseType
	MANAGED_ROMA   ShowInferServiceClusterResponseType
	DEDICATED      ShowInferServiceClusterResponseType
	DEDICATED_ROMA ShowInferServiceClusterResponseType
}

func GetShowInferServiceClusterResponseTypeEnum() ShowInferServiceClusterResponseTypeEnum {
	return ShowInferServiceClusterResponseTypeEnum{
		MANAGED: ShowInferServiceClusterResponseType{
			value: "MANAGED",
		},
		MANAGED_ROMA: ShowInferServiceClusterResponseType{
			value: "MANAGED_ROMA",
		},
		DEDICATED: ShowInferServiceClusterResponseType{
			value: "DEDICATED",
		},
		DEDICATED_ROMA: ShowInferServiceClusterResponseType{
			value: "DEDICATED_ROMA",
		},
	}
}

func (c ShowInferServiceClusterResponseType) Value() string {
	return c.value
}

func (c ShowInferServiceClusterResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInferServiceClusterResponseType) UnmarshalJSON(b []byte) error {
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
