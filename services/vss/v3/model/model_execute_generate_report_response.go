package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ExecuteGenerateReportResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *ExecuteGenerateReportResponseInfoCode `json:"info_code,omitempty" xml:"info_code"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty" xml:"info_description"`
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
