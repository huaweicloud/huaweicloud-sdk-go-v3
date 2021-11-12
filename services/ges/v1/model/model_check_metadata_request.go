package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CheckMetadataRequest struct {
	// 图actionId

	ActionId CheckMetadataRequestActionId `json:"action_id"`

	Body *CheckMetadataReq `json:"body,omitempty"`
}

func (o CheckMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckMetadataRequest struct{}"
	}

	return strings.Join([]string{"CheckMetadataRequest", string(data)}, " ")
}

type CheckMetadataRequestActionId struct {
	value string
}

type CheckMetadataRequestActionIdEnum struct {
	CHECK_SCHEMA CheckMetadataRequestActionId
}

func GetCheckMetadataRequestActionIdEnum() CheckMetadataRequestActionIdEnum {
	return CheckMetadataRequestActionIdEnum{
		CHECK_SCHEMA: CheckMetadataRequestActionId{
			value: "check-schema",
		},
	}
}

func (c CheckMetadataRequestActionId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckMetadataRequestActionId) UnmarshalJSON(b []byte) error {
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
