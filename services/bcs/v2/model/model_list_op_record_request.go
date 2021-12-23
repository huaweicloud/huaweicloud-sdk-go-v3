package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListOpRecordRequest struct {
	// 区块链ID

	BlockchainId *string `json:"blockchain_id,omitempty"`
	// 操作状态, 可选数值如下（括号中为该数值对应的操作类型含义）：waiting(等待执行)，processing（处理中），finished（操作完成，成功），failed（操作失败），stop（操作中断）

	OperationStatus *ListOpRecordRequestOperationStatus `json:"operation_status,omitempty"`
	// 资源类型, 可选数值如下（括号中为该数值对应的操作类型含义）：BCSSVC01（BCS变更操作），BCSSVC02（UGBaaS变更操作），PLUGIN01（RestAPI插件变更操作），PLUGIN02（TC3插件变更操作），PLUGIN03（轻节点插件变更操作）

	ResourceType *ListOpRecordRequestResourceType `json:"resource_type,omitempty"`
	// 操作类型,  可选数值如下（括号中为该数值对应的操作类型含义）：99（OpCreate）,00（OpDelete）,01（OpUpgrade）,91（OpUpgradeRb）,02（OpAddOrg）,03（OpScaleOrg）,04（OpJoinChannel）,05（OpJoinUnion）

	OperationType *ListOpRecordRequestOperationType `json:"operation_type,omitempty"`
	// 操作记录ID

	OperationId *string `json:"operation_id,omitempty"`
}

func (o ListOpRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOpRecordRequest struct{}"
	}

	return strings.Join([]string{"ListOpRecordRequest", string(data)}, " ")
}

type ListOpRecordRequestOperationStatus struct {
	value string
}

type ListOpRecordRequestOperationStatusEnum struct {
	WAITING    ListOpRecordRequestOperationStatus
	PROCESSING ListOpRecordRequestOperationStatus
	FINISHED   ListOpRecordRequestOperationStatus
	FAILED     ListOpRecordRequestOperationStatus
	STOP       ListOpRecordRequestOperationStatus
}

func GetListOpRecordRequestOperationStatusEnum() ListOpRecordRequestOperationStatusEnum {
	return ListOpRecordRequestOperationStatusEnum{
		WAITING: ListOpRecordRequestOperationStatus{
			value: "waiting",
		},
		PROCESSING: ListOpRecordRequestOperationStatus{
			value: "processing",
		},
		FINISHED: ListOpRecordRequestOperationStatus{
			value: "finished",
		},
		FAILED: ListOpRecordRequestOperationStatus{
			value: "failed",
		},
		STOP: ListOpRecordRequestOperationStatus{
			value: "stop",
		},
	}
}

func (c ListOpRecordRequestOperationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOpRecordRequestOperationStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListOpRecordRequestResourceType struct {
	value string
}

type ListOpRecordRequestResourceTypeEnum struct {
	BCSSVC01 ListOpRecordRequestResourceType
	BCSSVC02 ListOpRecordRequestResourceType
	PLUGIN01 ListOpRecordRequestResourceType
	PLUGIN02 ListOpRecordRequestResourceType
	PLUGIN03 ListOpRecordRequestResourceType
}

func GetListOpRecordRequestResourceTypeEnum() ListOpRecordRequestResourceTypeEnum {
	return ListOpRecordRequestResourceTypeEnum{
		BCSSVC01: ListOpRecordRequestResourceType{
			value: "BCSSVC01",
		},
		BCSSVC02: ListOpRecordRequestResourceType{
			value: "BCSSVC02",
		},
		PLUGIN01: ListOpRecordRequestResourceType{
			value: "PLUGIN01",
		},
		PLUGIN02: ListOpRecordRequestResourceType{
			value: "PLUGIN02",
		},
		PLUGIN03: ListOpRecordRequestResourceType{
			value: "PLUGIN03",
		},
	}
}

func (c ListOpRecordRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOpRecordRequestResourceType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type ListOpRecordRequestOperationType struct {
	value string
}

type ListOpRecordRequestOperationTypeEnum struct {
	E_99 ListOpRecordRequestOperationType
	E_00 ListOpRecordRequestOperationType
	E_01 ListOpRecordRequestOperationType
	E_91 ListOpRecordRequestOperationType
	E_02 ListOpRecordRequestOperationType
	E_03 ListOpRecordRequestOperationType
	E_04 ListOpRecordRequestOperationType
	E_05 ListOpRecordRequestOperationType
}

func GetListOpRecordRequestOperationTypeEnum() ListOpRecordRequestOperationTypeEnum {
	return ListOpRecordRequestOperationTypeEnum{
		E_99: ListOpRecordRequestOperationType{
			value: "99",
		},
		E_00: ListOpRecordRequestOperationType{
			value: "00",
		},
		E_01: ListOpRecordRequestOperationType{
			value: "01",
		},
		E_91: ListOpRecordRequestOperationType{
			value: "91",
		},
		E_02: ListOpRecordRequestOperationType{
			value: "02",
		},
		E_03: ListOpRecordRequestOperationType{
			value: "03",
		},
		E_04: ListOpRecordRequestOperationType{
			value: "04",
		},
		E_05: ListOpRecordRequestOperationType{
			value: "05",
		},
	}
}

func (c ListOpRecordRequestOperationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListOpRecordRequestOperationType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
