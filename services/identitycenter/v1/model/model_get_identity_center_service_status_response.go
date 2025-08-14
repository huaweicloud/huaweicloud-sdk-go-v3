package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GetIdentityCenterServiceStatusResponse Response Object
type GetIdentityCenterServiceStatusResponse struct {

	// IAM身份中心服务实例状态
	ServiceStatus *GetIdentityCenterServiceStatusResponseServiceStatus `json:"serviceStatus,omitempty"`

	// IAM身份中心服务实例状态呈现原因
	ServiceStatusReasons *[]string `json:"serviceStatusReasons,omitempty"`
	HttpStatusCode       int       `json:"-"`
}

func (o GetIdentityCenterServiceStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetIdentityCenterServiceStatusResponse struct{}"
	}

	return strings.Join([]string{"GetIdentityCenterServiceStatusResponse", string(data)}, " ")
}

type GetIdentityCenterServiceStatusResponseServiceStatus struct {
	value string
}

type GetIdentityCenterServiceStatusResponseServiceStatusEnum struct {
	ENABLED  GetIdentityCenterServiceStatusResponseServiceStatus
	DISABLED GetIdentityCenterServiceStatusResponseServiceStatus
}

func GetGetIdentityCenterServiceStatusResponseServiceStatusEnum() GetIdentityCenterServiceStatusResponseServiceStatusEnum {
	return GetIdentityCenterServiceStatusResponseServiceStatusEnum{
		ENABLED: GetIdentityCenterServiceStatusResponseServiceStatus{
			value: "ENABLED",
		},
		DISABLED: GetIdentityCenterServiceStatusResponseServiceStatus{
			value: "DISABLED",
		},
	}
}

func (c GetIdentityCenterServiceStatusResponseServiceStatus) Value() string {
	return c.value
}

func (c GetIdentityCenterServiceStatusResponseServiceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GetIdentityCenterServiceStatusResponseServiceStatus) UnmarshalJSON(b []byte) error {
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
