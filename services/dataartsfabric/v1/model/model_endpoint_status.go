package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EndpointStatus 状态。可选值如下： CREATING：创建中 RUNNING：运行中 CREATE_FAILED：创建失败 DELETING：删除中 DELETED：已删除 UPDATING: 更新中 UPDATE_FAILED:升级失败 DELETE_FAILED：创建失败 FROZEN:冻结 INACTIVE:未开通（公共EP）
type EndpointStatus struct {
	value string
}

type EndpointStatusEnum struct {
	CREATING      EndpointStatus
	RUNNING       EndpointStatus
	CREATE_FAILED EndpointStatus
	DELETING      EndpointStatus
	DELETED       EndpointStatus
	UPDATING      EndpointStatus
	UPDATE_FAILED EndpointStatus
	DELETE_FAILED EndpointStatus
	FROZEN        EndpointStatus
	INACTIVE      EndpointStatus
}

func GetEndpointStatusEnum() EndpointStatusEnum {
	return EndpointStatusEnum{
		CREATING: EndpointStatus{
			value: "CREATING",
		},
		RUNNING: EndpointStatus{
			value: "RUNNING",
		},
		CREATE_FAILED: EndpointStatus{
			value: "CREATE_FAILED",
		},
		DELETING: EndpointStatus{
			value: "DELETING",
		},
		DELETED: EndpointStatus{
			value: "DELETED",
		},
		UPDATING: EndpointStatus{
			value: "UPDATING",
		},
		UPDATE_FAILED: EndpointStatus{
			value: "UPDATE_FAILED",
		},
		DELETE_FAILED: EndpointStatus{
			value: "DELETE_FAILED",
		},
		FROZEN: EndpointStatus{
			value: "FROZEN",
		},
		INACTIVE: EndpointStatus{
			value: "INACTIVE",
		},
	}
}

func (c EndpointStatus) Value() string {
	return c.value
}

func (c EndpointStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EndpointStatus) UnmarshalJSON(b []byte) error {
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
