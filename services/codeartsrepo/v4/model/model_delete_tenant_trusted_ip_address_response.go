package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteTenantTrustedIpAddressResponse Response Object
type DeleteTenantTrustedIpAddressResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *DeleteTenantTrustedIpAddressResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o DeleteTenantTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTenantTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"DeleteTenantTrustedIpAddressResponse", string(data)}, " ")
}

type DeleteTenantTrustedIpAddressResponseStatus struct {
	value string
}

type DeleteTenantTrustedIpAddressResponseStatusEnum struct {
	SUCCESS DeleteTenantTrustedIpAddressResponseStatus
	FAIL    DeleteTenantTrustedIpAddressResponseStatus
}

func GetDeleteTenantTrustedIpAddressResponseStatusEnum() DeleteTenantTrustedIpAddressResponseStatusEnum {
	return DeleteTenantTrustedIpAddressResponseStatusEnum{
		SUCCESS: DeleteTenantTrustedIpAddressResponseStatus{
			value: "success",
		},
		FAIL: DeleteTenantTrustedIpAddressResponseStatus{
			value: "fail",
		},
	}
}

func (c DeleteTenantTrustedIpAddressResponseStatus) Value() string {
	return c.value
}

func (c DeleteTenantTrustedIpAddressResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTenantTrustedIpAddressResponseStatus) UnmarshalJSON(b []byte) error {
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
