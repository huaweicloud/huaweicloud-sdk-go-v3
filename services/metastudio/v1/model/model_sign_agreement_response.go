package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SignAgreementResponse Response Object
type SignAgreementResponse struct {

	// 当前服务协议版本
	CurrentVersion *string `json:"current_version,omitempty"`

	// 是否需要签署
	NeedToSign *bool `json:"need_to_sign,omitempty"`

	// 签署服务协议版本
	SignedVersion *string `json:"signed_version,omitempty"`

	// 协议签署时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"
	SignedTime *string `json:"signed_time,omitempty"`

	// 签署最后期限。格式遵循：RFC 3339 如\"2024-10-10T15:59:59Z\"
	ResignDeadline *string `json:"resign_deadline,omitempty"`

	// 重新签署过期处理方式:EXPIRE_AUTO_AGREE 自动签署协议,EXPIRE_STOP_SERVICE 强制签署协议
	ResignExpireProcessMode *SignAgreementResponseResignExpireProcessMode `json:"resign_expire_process_mode,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SignAgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignAgreementResponse struct{}"
	}

	return strings.Join([]string{"SignAgreementResponse", string(data)}, " ")
}

type SignAgreementResponseResignExpireProcessMode struct {
	value string
}

type SignAgreementResponseResignExpireProcessModeEnum struct {
	EXPIRE_AUTO_AGREE   SignAgreementResponseResignExpireProcessMode
	EXPIRE_STOP_SERVICE SignAgreementResponseResignExpireProcessMode
}

func GetSignAgreementResponseResignExpireProcessModeEnum() SignAgreementResponseResignExpireProcessModeEnum {
	return SignAgreementResponseResignExpireProcessModeEnum{
		EXPIRE_AUTO_AGREE: SignAgreementResponseResignExpireProcessMode{
			value: "EXPIRE_AUTO_AGREE",
		},
		EXPIRE_STOP_SERVICE: SignAgreementResponseResignExpireProcessMode{
			value: "EXPIRE_STOP_SERVICE",
		},
	}
}

func (c SignAgreementResponseResignExpireProcessMode) Value() string {
	return c.value
}

func (c SignAgreementResponseResignExpireProcessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignAgreementResponseResignExpireProcessMode) UnmarshalJSON(b []byte) error {
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
