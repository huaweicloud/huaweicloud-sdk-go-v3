package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExecuteGenerateReportResponse Response Object
type ExecuteGenerateReportResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *ExecuteGenerateReportResponseInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ExecuteGenerateReportResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteGenerateReportResponse struct{}"
	}

	return strings.Join([]string{"ExecuteGenerateReportResponse", string(data)}, " ")
}

type ExecuteGenerateReportResponseInfoCode struct {
	value string
}

type ExecuteGenerateReportResponseInfoCodeEnum struct {
	SUCCESS ExecuteGenerateReportResponseInfoCode
	FAILURE ExecuteGenerateReportResponseInfoCode
}

func GetExecuteGenerateReportResponseInfoCodeEnum() ExecuteGenerateReportResponseInfoCodeEnum {
	return ExecuteGenerateReportResponseInfoCodeEnum{
		SUCCESS: ExecuteGenerateReportResponseInfoCode{
			value: "success",
		},
		FAILURE: ExecuteGenerateReportResponseInfoCode{
			value: "failure",
		},
	}
}

func (c ExecuteGenerateReportResponseInfoCode) Value() string {
	return c.value
}

func (c ExecuteGenerateReportResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteGenerateReportResponseInfoCode) UnmarshalJSON(b []byte) error {
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
