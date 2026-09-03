package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyStatus 策略的当前状态。CREATING - 策略正在创建中；ACTIVE - 策略已激活并可用；UPDATING - 策略正在更新中；DELETING - 策略正在删除中；CREATE_FAILED - 策略创建失败；UPDATE_FAILED - 策略更新失败；DELETE_FAILED - 策略删除失败。
type PolicyStatus struct {
	value string
}

type PolicyStatusEnum struct {
	CREATING      PolicyStatus
	ACTIVE        PolicyStatus
	UPDATING      PolicyStatus
	DELETING      PolicyStatus
	CREATE_FAILED PolicyStatus
	UPDATE_FAILED PolicyStatus
	DELETE_FAILED PolicyStatus
}

func GetPolicyStatusEnum() PolicyStatusEnum {
	return PolicyStatusEnum{
		CREATING: PolicyStatus{
			value: "CREATING",
		},
		ACTIVE: PolicyStatus{
			value: "ACTIVE",
		},
		UPDATING: PolicyStatus{
			value: "UPDATING",
		},
		DELETING: PolicyStatus{
			value: "DELETING",
		},
		CREATE_FAILED: PolicyStatus{
			value: "CREATE_FAILED",
		},
		UPDATE_FAILED: PolicyStatus{
			value: "UPDATE_FAILED",
		},
		DELETE_FAILED: PolicyStatus{
			value: "DELETE_FAILED",
		},
	}
}

func (c PolicyStatus) Value() string {
	return c.value
}

func (c PolicyStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyStatus) UnmarshalJSON(b []byte) error {
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
