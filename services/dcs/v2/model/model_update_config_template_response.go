package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateConfigTemplateResponse Response Object
type UpdateConfigTemplateResponse struct {

	// 修改自定义模板的结果，success表示成功，如果失败则会返回对应的报错信息
	Result         *UpdateConfigTemplateResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o UpdateConfigTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConfigTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateConfigTemplateResponse", string(data)}, " ")
}

type UpdateConfigTemplateResponseResult struct {
	value string
}

type UpdateConfigTemplateResponseResultEnum struct {
	SUCCESS UpdateConfigTemplateResponseResult
}

func GetUpdateConfigTemplateResponseResultEnum() UpdateConfigTemplateResponseResultEnum {
	return UpdateConfigTemplateResponseResultEnum{
		SUCCESS: UpdateConfigTemplateResponseResult{
			value: "success",
		},
	}
}

func (c UpdateConfigTemplateResponseResult) Value() string {
	return c.value
}

func (c UpdateConfigTemplateResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateConfigTemplateResponseResult) UnmarshalJSON(b []byte) error {
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
