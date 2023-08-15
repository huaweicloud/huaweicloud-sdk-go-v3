package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfirmImageUploadReq
type ConfirmImageUploadReq struct {

	// 水印配置模板id。
	Id string `json:"id"`

	// 水印上传状态。
	Status ConfirmImageUploadReqStatus `json:"status"`
}

func (o ConfirmImageUploadReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfirmImageUploadReq struct{}"
	}

	return strings.Join([]string{"ConfirmImageUploadReq", string(data)}, " ")
}

type ConfirmImageUploadReqStatus struct {
	value string
}

type ConfirmImageUploadReqStatusEnum struct {
	SUCCEED ConfirmImageUploadReqStatus
	FAILED  ConfirmImageUploadReqStatus
}

func GetConfirmImageUploadReqStatusEnum() ConfirmImageUploadReqStatusEnum {
	return ConfirmImageUploadReqStatusEnum{
		SUCCEED: ConfirmImageUploadReqStatus{
			value: "SUCCEED",
		},
		FAILED: ConfirmImageUploadReqStatus{
			value: "FAILED",
		},
	}
}

func (c ConfirmImageUploadReqStatus) Value() string {
	return c.value
}

func (c ConfirmImageUploadReqStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfirmImageUploadReqStatus) UnmarshalJSON(b []byte) error {
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
