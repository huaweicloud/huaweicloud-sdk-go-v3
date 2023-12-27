package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CallBackConfig struct {

	// 回调URL。  回调请求body为json格式，带参数如下：  result: SUCCEED或FAILED  asset_id: 资产ID  job_id: 任务
	CallbackUrl string `json:"callback_url"`

	// 认证类型。 * NONE。URL中自带认证。 * MSS_A。HMACSHA256签名模式，在URL中追加参数:secret,time_stamp。取值方式：secret=hmac_sha256(key, URI（callback_url）+ time_stamp)&time_stamp=hex(timestamp)
	AuthType *CallBackConfigAuthType `json:"auth_type,omitempty"`

	// 密钥Key
	Key *string `json:"key,omitempty"`
}

func (o CallBackConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CallBackConfig struct{}"
	}

	return strings.Join([]string{"CallBackConfig", string(data)}, " ")
}

type CallBackConfigAuthType struct {
	value string
}

type CallBackConfigAuthTypeEnum struct {
	NONE  CallBackConfigAuthType
	MSS_A CallBackConfigAuthType
}

func GetCallBackConfigAuthTypeEnum() CallBackConfigAuthTypeEnum {
	return CallBackConfigAuthTypeEnum{
		NONE: CallBackConfigAuthType{
			value: "NONE",
		},
		MSS_A: CallBackConfigAuthType{
			value: "MSS_A",
		},
	}
}

func (c CallBackConfigAuthType) Value() string {
	return c.value
}

func (c CallBackConfigAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CallBackConfigAuthType) UnmarshalJSON(b []byte) error {
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
