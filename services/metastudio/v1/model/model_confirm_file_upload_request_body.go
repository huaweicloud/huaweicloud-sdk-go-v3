package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ConfirmFileUploadRequestBody struct {

	// **参数解释**： 文件上传状态。 **约束限制**： 不涉及 **取值范围**： * CREATED：上传完成 * FAILED：上传失败 * CANCELLED：取消上传  **默认取值**： 不涉及
	State ConfirmFileUploadRequestBodyState `json:"state"`
}

func (o ConfirmFileUploadRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmFileUploadRequestBody struct{}"
	}

	return strings.Join([]string{"ConfirmFileUploadRequestBody", string(data)}, " ")
}

type ConfirmFileUploadRequestBodyState struct {
	value string
}

type ConfirmFileUploadRequestBodyStateEnum struct {
	CREATED   ConfirmFileUploadRequestBodyState
	FAILED    ConfirmFileUploadRequestBodyState
	CANCELLED ConfirmFileUploadRequestBodyState
}

func GetConfirmFileUploadRequestBodyStateEnum() ConfirmFileUploadRequestBodyStateEnum {
	return ConfirmFileUploadRequestBodyStateEnum{
		CREATED: ConfirmFileUploadRequestBodyState{
			value: "CREATED",
		},
		FAILED: ConfirmFileUploadRequestBodyState{
			value: "FAILED",
		},
		CANCELLED: ConfirmFileUploadRequestBodyState{
			value: "CANCELLED",
		},
	}
}

func (c ConfirmFileUploadRequestBodyState) Value() string {
	return c.value
}

func (c ConfirmFileUploadRequestBodyState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmFileUploadRequestBodyState) UnmarshalJSON(b []byte) error {
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
