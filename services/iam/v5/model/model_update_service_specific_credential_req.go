package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateServiceSpecificCredentialReq 更新服务专属凭证的请求体。
type UpdateServiceSpecificCredentialReq struct {

	// 凭证状态。不支持用户主动设置为 Expired（Expired 仅由系统在凭证过期时自动设置）。status和description至少指定一个。
	Status *UpdateServiceSpecificCredentialReqStatus `json:"status,omitempty"`

	// 要设置的凭证描述。status和description至少指定一个，长度0-255，正则限制为^[^@#%&<>\\\\\\$\\^\\*]*$
	Description *string `json:"description,omitempty"`
}

func (o UpdateServiceSpecificCredentialReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateServiceSpecificCredentialReq struct{}"
	}

	return strings.Join([]string{"UpdateServiceSpecificCredentialReq", string(data)}, " ")
}

type UpdateServiceSpecificCredentialReqStatus struct {
	value string
}

type UpdateServiceSpecificCredentialReqStatusEnum struct {
	ACTIVE   UpdateServiceSpecificCredentialReqStatus
	INACTIVE UpdateServiceSpecificCredentialReqStatus
}

func GetUpdateServiceSpecificCredentialReqStatusEnum() UpdateServiceSpecificCredentialReqStatusEnum {
	return UpdateServiceSpecificCredentialReqStatusEnum{
		ACTIVE: UpdateServiceSpecificCredentialReqStatus{
			value: "active",
		},
		INACTIVE: UpdateServiceSpecificCredentialReqStatus{
			value: "inactive",
		},
	}
}

func (c UpdateServiceSpecificCredentialReqStatus) Value() string {
	return c.value
}

func (c UpdateServiceSpecificCredentialReqStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateServiceSpecificCredentialReqStatus) UnmarshalJSON(b []byte) error {
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
