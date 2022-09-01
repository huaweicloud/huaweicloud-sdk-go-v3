package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type DeleteDomainsResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *DeleteDomainsResponseInfoCode `json:"info_code,omitempty" xml:"info_code"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty" xml:"info_description"`
	HttpStatusCode  int     `json:"-"`
}

func (o DeleteDomainsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDomainsResponse struct{}"
	}

	return strings.Join([]string{"DeleteDomainsResponse", string(data)}, " ")
}

type DeleteDomainsResponseInfoCode struct {
	value string
}

type DeleteDomainsResponseInfoCodeEnum struct {
	SUCCESS DeleteDomainsResponseInfoCode
	FAILURE DeleteDomainsResponseInfoCode
}

func GetDeleteDomainsResponseInfoCodeEnum() DeleteDomainsResponseInfoCodeEnum {
	return DeleteDomainsResponseInfoCodeEnum{
		SUCCESS: DeleteDomainsResponseInfoCode{
			value: "success",
		},
		FAILURE: DeleteDomainsResponseInfoCode{
			value: "failure",
		},
	}
}

func (c DeleteDomainsResponseInfoCode) Value() string {
	return c.value
}

func (c DeleteDomainsResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteDomainsResponseInfoCode) UnmarshalJSON(b []byte) error {
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
