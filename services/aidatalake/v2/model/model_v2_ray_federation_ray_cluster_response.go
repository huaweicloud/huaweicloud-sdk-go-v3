package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// V2RayFederationRayClusterResponse **参数解释**：Ray联邦端点中RayCluster端点的信息（响应）。
type V2RayFederationRayClusterResponse struct {

	// **参数解释**：RayCluster端点ID。 **取值范围**：长度为32~36的英文字符、数字和中划线的组合。
	Id *string `json:"id,omitempty"`

	// **参数解释**：RayCluster端点名称。 **取值范围**：支持小写字母、数字、中划线，必须字母开头、字母或数字结尾，长度为1~63个字符。
	Name *string `json:"name,omitempty"`

	// **参数解释**：RayCluster端点状态。 **取值范围**：   - CREATING：创建中。   - RUNNING：运行中。   - CREATE_FAILED：创建失败。   - DELETING：删除中。   - DELETED：已删除。   - UPDATING：更新中。   - UPDATE_FAILED：更新失败。   - DELETE_FAILED：删除失败。   - FROZEN：冻结。   - INACTIVE：未激活。   - ON_DEMAND_FROZEN：按需冻结。   - ABNORMAL：异常。
	Status *V2RayFederationRayClusterResponseStatus `json:"status,omitempty"`
}

func (o V2RayFederationRayClusterResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "V2RayFederationRayClusterResponse struct{}"
	}

	return strings.Join([]string{"V2RayFederationRayClusterResponse", string(data)}, " ")
}

type V2RayFederationRayClusterResponseStatus struct {
	value string
}

type V2RayFederationRayClusterResponseStatusEnum struct {
	CREATING         V2RayFederationRayClusterResponseStatus
	RUNNING          V2RayFederationRayClusterResponseStatus
	CREATE_FAILED    V2RayFederationRayClusterResponseStatus
	DELETING         V2RayFederationRayClusterResponseStatus
	DELETED          V2RayFederationRayClusterResponseStatus
	UPDATING         V2RayFederationRayClusterResponseStatus
	UPDATE_FAILED    V2RayFederationRayClusterResponseStatus
	DELETE_FAILED    V2RayFederationRayClusterResponseStatus
	FROZEN           V2RayFederationRayClusterResponseStatus
	INACTIVE         V2RayFederationRayClusterResponseStatus
	ON_DEMAND_FROZEN V2RayFederationRayClusterResponseStatus
	ABNORMAL         V2RayFederationRayClusterResponseStatus
}

func GetV2RayFederationRayClusterResponseStatusEnum() V2RayFederationRayClusterResponseStatusEnum {
	return V2RayFederationRayClusterResponseStatusEnum{
		CREATING: V2RayFederationRayClusterResponseStatus{
			value: "CREATING",
		},
		RUNNING: V2RayFederationRayClusterResponseStatus{
			value: "RUNNING",
		},
		CREATE_FAILED: V2RayFederationRayClusterResponseStatus{
			value: "CREATE_FAILED",
		},
		DELETING: V2RayFederationRayClusterResponseStatus{
			value: "DELETING",
		},
		DELETED: V2RayFederationRayClusterResponseStatus{
			value: "DELETED",
		},
		UPDATING: V2RayFederationRayClusterResponseStatus{
			value: "UPDATING",
		},
		UPDATE_FAILED: V2RayFederationRayClusterResponseStatus{
			value: "UPDATE_FAILED",
		},
		DELETE_FAILED: V2RayFederationRayClusterResponseStatus{
			value: "DELETE_FAILED",
		},
		FROZEN: V2RayFederationRayClusterResponseStatus{
			value: "FROZEN",
		},
		INACTIVE: V2RayFederationRayClusterResponseStatus{
			value: "INACTIVE",
		},
		ON_DEMAND_FROZEN: V2RayFederationRayClusterResponseStatus{
			value: "ON_DEMAND_FROZEN",
		},
		ABNORMAL: V2RayFederationRayClusterResponseStatus{
			value: "ABNORMAL",
		},
	}
}

func (c V2RayFederationRayClusterResponseStatus) Value() string {
	return c.value
}

func (c V2RayFederationRayClusterResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *V2RayFederationRayClusterResponseStatus) UnmarshalJSON(b []byte) error {
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
