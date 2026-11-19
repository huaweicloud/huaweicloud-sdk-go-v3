package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTrustedKeySpaceRequestBody 创建密钥空间请求体
type CreateTrustedKeySpaceRequestBody struct {

	// **参数解释：** 密钥空间名称 **约束限制：** 不涉及 **取值范围：** 1-255 **默认取值：** 不涉及
	KeyspaceName string `json:"keyspace_name"`

	// **参数解释：** 密钥空间类型 **约束限制：** 不涉及 **取值范围：** - 1 默认密钥空间 - 2 专属加密集群密钥空间 - 3 外部密钥存储密钥空间 **默认取值：** 1
	Type *CreateTrustedKeySpaceRequestBodyType `json:"type,omitempty"`

	// **参数解释：** 密钥空间关联的密钥库ID,仅type为2或3时必填，关联的密钥库ID需要和类型匹配 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	KeystoreId *string `json:"keystore_id,omitempty"`

	// **参数解释：** 密钥空间描述信息 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Description *string `json:"description,omitempty"`
}

func (o CreateTrustedKeySpaceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTrustedKeySpaceRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTrustedKeySpaceRequestBody", string(data)}, " ")
}

type CreateTrustedKeySpaceRequestBodyType struct {
	value string
}

type CreateTrustedKeySpaceRequestBodyTypeEnum struct {
	E_1 CreateTrustedKeySpaceRequestBodyType
	E_2 CreateTrustedKeySpaceRequestBodyType
	E_3 CreateTrustedKeySpaceRequestBodyType
}

func GetCreateTrustedKeySpaceRequestBodyTypeEnum() CreateTrustedKeySpaceRequestBodyTypeEnum {
	return CreateTrustedKeySpaceRequestBodyTypeEnum{
		E_1: CreateTrustedKeySpaceRequestBodyType{
			value: "1",
		},
		E_2: CreateTrustedKeySpaceRequestBodyType{
			value: "2",
		},
		E_3: CreateTrustedKeySpaceRequestBodyType{
			value: "3",
		},
	}
}

func (c CreateTrustedKeySpaceRequestBodyType) Value() string {
	return c.value
}

func (c CreateTrustedKeySpaceRequestBodyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTrustedKeySpaceRequestBodyType) UnmarshalJSON(b []byte) error {
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
