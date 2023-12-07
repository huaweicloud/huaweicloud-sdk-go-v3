package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApplyDetailRequest Request Object
type ShowApplyDetailRequest struct {

	// 工作空间id
	Workspace string `json:"workspace"`

	// dlm版本类型
	DlmType *ShowApplyDetailRequestDlmType `json:"Dlm-Type,omitempty"`

	// 资源类型
	ContentType string `json:"Content-Type"`

	// 审核信息id
	ApplyId string `json:"apply_id"`
}

func (o ShowApplyDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApplyDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowApplyDetailRequest", string(data)}, " ")
}

type ShowApplyDetailRequestDlmType struct {
	value string
}

type ShowApplyDetailRequestDlmTypeEnum struct {
	SHARED    ShowApplyDetailRequestDlmType
	EXCLUSIVE ShowApplyDetailRequestDlmType
}

func GetShowApplyDetailRequestDlmTypeEnum() ShowApplyDetailRequestDlmTypeEnum {
	return ShowApplyDetailRequestDlmTypeEnum{
		SHARED: ShowApplyDetailRequestDlmType{
			value: "SHARED",
		},
		EXCLUSIVE: ShowApplyDetailRequestDlmType{
			value: "EXCLUSIVE",
		},
	}
}

func (c ShowApplyDetailRequestDlmType) Value() string {
	return c.value
}

func (c ShowApplyDetailRequestDlmType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApplyDetailRequestDlmType) UnmarshalJSON(b []byte) error {
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
