package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CentralNetworkElementChange 中心网络策略变化。
type CentralNetworkElementChange struct {

	// 实例状态。 - CreateCentralNetworkPlane: 新增中心网络平面 - DeleteCentralNetworkPlane: 移除中心网络平面 - UpdateCentralNetworkPlane: 更新中心网络平面 - CreateCentralNetworkErInstance: 新增中心网络ER实例 - DeleteCentralNetworkErInstance: 移除中心网络ER实例 - CreateCentralNetworkErConnection: 新增中心网络ER连接 - DeleteCentralNetworkErConnection: 移除中心网络ER连接 - CreateCentralNetworkErTable: 新增中心网络ER路由表 - DeleteCentralNetworkErTable: 移除中心网络ER路由表
	OperationId CentralNetworkElementChangeOperationId `json:"operation_id"`
}

func (o CentralNetworkElementChange) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CentralNetworkElementChange struct{}"
	}

	return strings.Join([]string{"CentralNetworkElementChange", string(data)}, " ")
}

type CentralNetworkElementChangeOperationId struct {
	value string
}

type CentralNetworkElementChangeOperationIdEnum struct {
	CREATE_CENTRAL_NETWORK_PLANE         CentralNetworkElementChangeOperationId
	DELETE_CENTRAL_NETWORK_PLANE         CentralNetworkElementChangeOperationId
	UPDATE_CENTRAL_NETWORK_PLANE         CentralNetworkElementChangeOperationId
	CREATE_CENTRAL_NETWORK_ER_INSTANCE   CentralNetworkElementChangeOperationId
	DELETE_CENTRAL_NETWORK_ER_INSTANCE   CentralNetworkElementChangeOperationId
	CREATE_CENTRAL_NETWORK_ER_CONNECTION CentralNetworkElementChangeOperationId
	DELETE_CENTRAL_NETWORK_ER_CONNECTION CentralNetworkElementChangeOperationId
	CREATE_CENTRAL_NETWORK_ER_TABLE      CentralNetworkElementChangeOperationId
	DELETE_CENTRAL_NETWORK_ER_TABLE      CentralNetworkElementChangeOperationId
}

func GetCentralNetworkElementChangeOperationIdEnum() CentralNetworkElementChangeOperationIdEnum {
	return CentralNetworkElementChangeOperationIdEnum{
		CREATE_CENTRAL_NETWORK_PLANE: CentralNetworkElementChangeOperationId{
			value: "CreateCentralNetworkPlane",
		},
		DELETE_CENTRAL_NETWORK_PLANE: CentralNetworkElementChangeOperationId{
			value: "DeleteCentralNetworkPlane",
		},
		UPDATE_CENTRAL_NETWORK_PLANE: CentralNetworkElementChangeOperationId{
			value: "UpdateCentralNetworkPlane",
		},
		CREATE_CENTRAL_NETWORK_ER_INSTANCE: CentralNetworkElementChangeOperationId{
			value: "CreateCentralNetworkErInstance",
		},
		DELETE_CENTRAL_NETWORK_ER_INSTANCE: CentralNetworkElementChangeOperationId{
			value: "DeleteCentralNetworkErInstance",
		},
		CREATE_CENTRAL_NETWORK_ER_CONNECTION: CentralNetworkElementChangeOperationId{
			value: "CreateCentralNetworkErConnection",
		},
		DELETE_CENTRAL_NETWORK_ER_CONNECTION: CentralNetworkElementChangeOperationId{
			value: "DeleteCentralNetworkErConnection",
		},
		CREATE_CENTRAL_NETWORK_ER_TABLE: CentralNetworkElementChangeOperationId{
			value: "CreateCentralNetworkErTable",
		},
		DELETE_CENTRAL_NETWORK_ER_TABLE: CentralNetworkElementChangeOperationId{
			value: "DeleteCentralNetworkErTable",
		},
	}
}

func (c CentralNetworkElementChangeOperationId) Value() string {
	return c.value
}

func (c CentralNetworkElementChangeOperationId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CentralNetworkElementChangeOperationId) UnmarshalJSON(b []byte) error {
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
