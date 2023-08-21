package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateEdgeDDoSDomainsRequestBody struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 防护开关（0:关，1:开）
	ProtectedSwitch UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch `json:"protected_switch"`
}

func (o UpdateEdgeDDoSDomainsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeDDoSDomainsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateEdgeDDoSDomainsRequestBody", string(data)}, " ")
}

type UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch struct {
	value int32
}

type UpdateEdgeDDoSDomainsRequestBodyProtectedSwitchEnum struct {
	E_0 UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch
	E_1 UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch
}

func GetUpdateEdgeDDoSDomainsRequestBodyProtectedSwitchEnum() UpdateEdgeDDoSDomainsRequestBodyProtectedSwitchEnum {
	return UpdateEdgeDDoSDomainsRequestBodyProtectedSwitchEnum{
		E_0: UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch{
			value: 0,
		}, E_1: UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch{
			value: 1,
		},
	}
}

func (c UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch) Value() int32 {
	return c.value
}

func (c UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEdgeDDoSDomainsRequestBodyProtectedSwitch) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
