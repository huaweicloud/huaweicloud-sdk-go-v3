package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RayFederationRayClusterResponse **参数解释**：Ray联邦端点中RayCluster端点的信息（响应）。 **约束限制**：不涉及。
type RayFederationRayClusterResponse struct {

	// **参数解释**：RayCluster端点ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：RayCluster端点名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**：RayCluster端点状态。 **取值范围**：   - CREATING：创建中。   - RUNNING：运行中。   - CREATE_FAILED：创建失败。   - DELETING：删除中。   - DELETED：已删除。   - UPDATING：更新中。   - UPDATE_FAILED：更新失败。   - DELETE_FAILED：删除失败。   - FROZEN：冻结。   - INACTIVE：未激活。   - ON_DEMAND_FROZEN：按需冻结。   - ABNORMAL：异常。
	Status *RayFederationRayClusterResponseStatus `json:"status,omitempty"`
}

func (o RayFederationRayClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayFederationRayClusterResponse struct{}"
	}

	return strings.Join([]string{"RayFederationRayClusterResponse", string(data)}, " ")
}

type RayFederationRayClusterResponseStatus struct {
	value string
}

type RayFederationRayClusterResponseStatusEnum struct {
	CREATING         RayFederationRayClusterResponseStatus
	RUNNING          RayFederationRayClusterResponseStatus
	CREATE_FAILED    RayFederationRayClusterResponseStatus
	DELETING         RayFederationRayClusterResponseStatus
	DELETED          RayFederationRayClusterResponseStatus
	UPDATING         RayFederationRayClusterResponseStatus
	UPDATE_FAILED    RayFederationRayClusterResponseStatus
	DELETE_FAILED    RayFederationRayClusterResponseStatus
	FROZEN           RayFederationRayClusterResponseStatus
	INACTIVE         RayFederationRayClusterResponseStatus
	ON_DEMAND_FROZEN RayFederationRayClusterResponseStatus
	ABNORMAL         RayFederationRayClusterResponseStatus
}

func GetRayFederationRayClusterResponseStatusEnum() RayFederationRayClusterResponseStatusEnum {
	return RayFederationRayClusterResponseStatusEnum{
		CREATING: RayFederationRayClusterResponseStatus{
			value: "CREATING",
		},
		RUNNING: RayFederationRayClusterResponseStatus{
			value: "RUNNING",
		},
		CREATE_FAILED: RayFederationRayClusterResponseStatus{
			value: "CREATE_FAILED",
		},
		DELETING: RayFederationRayClusterResponseStatus{
			value: "DELETING",
		},
		DELETED: RayFederationRayClusterResponseStatus{
			value: "DELETED",
		},
		UPDATING: RayFederationRayClusterResponseStatus{
			value: "UPDATING",
		},
		UPDATE_FAILED: RayFederationRayClusterResponseStatus{
			value: "UPDATE_FAILED",
		},
		DELETE_FAILED: RayFederationRayClusterResponseStatus{
			value: "DELETE_FAILED",
		},
		FROZEN: RayFederationRayClusterResponseStatus{
			value: "FROZEN",
		},
		INACTIVE: RayFederationRayClusterResponseStatus{
			value: "INACTIVE",
		},
		ON_DEMAND_FROZEN: RayFederationRayClusterResponseStatus{
			value: "ON_DEMAND_FROZEN",
		},
		ABNORMAL: RayFederationRayClusterResponseStatus{
			value: "ABNORMAL",
		},
	}
}

func (c RayFederationRayClusterResponseStatus) Value() string {
	return c.value
}

func (c RayFederationRayClusterResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RayFederationRayClusterResponseStatus) UnmarshalJSON(b []byte) error {
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
