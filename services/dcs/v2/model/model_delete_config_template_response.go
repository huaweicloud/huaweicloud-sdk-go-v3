package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteConfigTemplateResponse Response Object
type DeleteConfigTemplateResponse struct {

	// 删除自定义模板的结果，success表示成功，如果失败则会返回对应的报错信息
	Result         *DeleteConfigTemplateResponseResult `json:"result,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o DeleteConfigTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConfigTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteConfigTemplateResponse", string(data)}, " ")
}

type DeleteConfigTemplateResponseResult struct {
	value string
}

type DeleteConfigTemplateResponseResultEnum struct {
	SUCCESS DeleteConfigTemplateResponseResult
}

func GetDeleteConfigTemplateResponseResultEnum() DeleteConfigTemplateResponseResultEnum {
	return DeleteConfigTemplateResponseResultEnum{
		SUCCESS: DeleteConfigTemplateResponseResult{
			value: "success",
		},
	}
}

func (c DeleteConfigTemplateResponseResult) Value() string {
	return c.value
}

func (c DeleteConfigTemplateResponseResult) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteConfigTemplateResponseResult) UnmarshalJSON(b []byte) error {
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
