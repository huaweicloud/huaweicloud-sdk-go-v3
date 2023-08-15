package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateDomainSettingsResponse Response Object
type UpdateDomainSettingsResponse struct {

	// 状态码:   * success - 成功   * failure - 失败
	InfoCode *UpdateDomainSettingsResponseInfoCode `json:"info_code,omitempty"`

	// 返回的提示信息
	InfoDescription *string `json:"info_description,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o UpdateDomainSettingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainSettingsResponse struct{}"
	}

	return strings.Join([]string{"UpdateDomainSettingsResponse", string(data)}, " ")
}

type UpdateDomainSettingsResponseInfoCode struct {
	value string
}

type UpdateDomainSettingsResponseInfoCodeEnum struct {
	SUCCESS UpdateDomainSettingsResponseInfoCode
	FAILURE UpdateDomainSettingsResponseInfoCode
}

func GetUpdateDomainSettingsResponseInfoCodeEnum() UpdateDomainSettingsResponseInfoCodeEnum {
	return UpdateDomainSettingsResponseInfoCodeEnum{
		SUCCESS: UpdateDomainSettingsResponseInfoCode{
			value: "success",
		},
		FAILURE: UpdateDomainSettingsResponseInfoCode{
			value: "failure",
		},
	}
}

func (c UpdateDomainSettingsResponseInfoCode) Value() string {
	return c.value
}

func (c UpdateDomainSettingsResponseInfoCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateDomainSettingsResponseInfoCode) UnmarshalJSON(b []byte) error {
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
