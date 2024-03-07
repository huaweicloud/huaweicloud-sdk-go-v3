package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessPreviewStatus 访问预览的状态。
type AccessPreviewStatus struct {
	value string
}

type AccessPreviewStatusEnum struct {
	CREATING  AccessPreviewStatus
	COMPLETED AccessPreviewStatus
	FAILED    AccessPreviewStatus
}

func GetAccessPreviewStatusEnum() AccessPreviewStatusEnum {
	return AccessPreviewStatusEnum{
		CREATING: AccessPreviewStatus{
			value: "creating",
		},
		COMPLETED: AccessPreviewStatus{
			value: "completed",
		},
		FAILED: AccessPreviewStatus{
			value: "failed",
		},
	}
}

func (c AccessPreviewStatus) Value() string {
	return c.value
}

func (c AccessPreviewStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessPreviewStatus) UnmarshalJSON(b []byte) error {
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
