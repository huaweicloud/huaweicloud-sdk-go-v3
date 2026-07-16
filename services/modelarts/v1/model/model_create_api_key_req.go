package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateApiKeyReq 创建api-key请求体
type CreateApiKeyReq struct {

	// **参数解释：** api-key名称，用户在[创建API_KEY](CreateInferApiKey.xml)时自定义。 **约束限制：** api-key在删除之前名字不能重复。 **取值范围：** 支持1-64个字符，可以包含字母、汉字、数字、连字符和下划线。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** api-key描述。 **约束限制：** 不涉及。 **取值范围：** 长度不可以超过256，不能包含感叹号，大于号，小于号，等号，与，单引号，双引号。 **默认取值：** 默认为空。
	Description *string `json:"description,omitempty"`

	// **参数解释：** api-key生效范围。 **约束限制：** 不涉及。 **取值范围：** - USER：表示生效范围为用户级别，可以访问该用户创建的所有在线服务。 - SERVICE：表示生效范围为单个服务，可以访问绑定该api-key的在线服务。 **默认取值：** 不涉及。
	Scope CreateApiKeyReqScope `json:"scope"`
}

func (o CreateApiKeyReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApiKeyReq struct{}"
	}

	return strings.Join([]string{"CreateApiKeyReq", string(data)}, " ")
}

type CreateApiKeyReqScope struct {
	value string
}

type CreateApiKeyReqScopeEnum struct {
	USER    CreateApiKeyReqScope
	SERVICE CreateApiKeyReqScope
}

func GetCreateApiKeyReqScopeEnum() CreateApiKeyReqScopeEnum {
	return CreateApiKeyReqScopeEnum{
		USER: CreateApiKeyReqScope{
			value: "USER",
		},
		SERVICE: CreateApiKeyReqScope{
			value: "SERVICE",
		},
	}
}

func (c CreateApiKeyReqScope) Value() string {
	return c.value
}

func (c CreateApiKeyReqScope) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateApiKeyReqScope) UnmarshalJSON(b []byte) error {
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
