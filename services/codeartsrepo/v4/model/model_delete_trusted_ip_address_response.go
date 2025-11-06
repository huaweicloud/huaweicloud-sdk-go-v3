package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteTrustedIpAddressResponse Response Object
type DeleteTrustedIpAddressResponse struct {

	// **参数解释：** 状态码。 - success，表示接口请求成功。 - fail，表示接口请求失败。
	Status         *DeleteTrustedIpAddressResponseStatus `json:"status,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o DeleteTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"DeleteTrustedIpAddressResponse", string(data)}, " ")
}

type DeleteTrustedIpAddressResponseStatus struct {
	value string
}

type DeleteTrustedIpAddressResponseStatusEnum struct {
	SUCCESS DeleteTrustedIpAddressResponseStatus
	FAIL    DeleteTrustedIpAddressResponseStatus
}

func GetDeleteTrustedIpAddressResponseStatusEnum() DeleteTrustedIpAddressResponseStatusEnum {
	return DeleteTrustedIpAddressResponseStatusEnum{
		SUCCESS: DeleteTrustedIpAddressResponseStatus{
			value: "success",
		},
		FAIL: DeleteTrustedIpAddressResponseStatus{
			value: "fail",
		},
	}
}

func (c DeleteTrustedIpAddressResponseStatus) Value() string {
	return c.value
}

func (c DeleteTrustedIpAddressResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteTrustedIpAddressResponseStatus) UnmarshalJSON(b []byte) error {
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
