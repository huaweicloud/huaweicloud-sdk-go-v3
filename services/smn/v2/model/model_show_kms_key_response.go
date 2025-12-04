package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowKmsKeyResponse Response Object
type ShowKmsKeyResponse struct {

	// 在DEW服务上创建的用户主密钥对应的密钥ID，具体参考在DEW服务上创建密钥章节。
	KeyId *string `json:"key_id,omitempty"`

	// 密钥的状态字段，为枚举类型。 - ENABLED状态为密钥启用状态，此时发布的消息均使用该主密钥创建的数据密钥进行加解密。 - TO_BE_ACTIVATED状态为密钥待激活状态，当密钥未激活时，发布主题消息会返回失败，请前往DEW服务对密钥进行操作。 - DISABLED状态为密钥禁用状态，当密钥已被禁用时，发布主题消息会返回失败，请前往DEW服务对密钥进行操作。 - PENDING_DELETION状态为密钥计划删除状态，当密钥已被计划删除时，发布主题消息仍能正常使用该密钥。 - PENDING_IMPORT状态为密钥计划导入状态，当密钥计划导入时，发布主题消息会返回失败，请前往DEW服务对密钥进行操作。 - DELETED状态为密钥已删除状态，当密钥已删除后，SMN无法从DEW服务处获取有效的密钥，此时发布主题消息会返回失败，请重新在主题下配置可用的密钥。
	Status *ShowKmsKeyResponseStatus `json:"status,omitempty"`

	// 密钥对应的资源ID，该ID由SMN服务生成。
	Id *string `json:"id,omitempty"`

	// 请求的唯一标识ID。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKmsKeyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKmsKeyResponse struct{}"
	}

	return strings.Join([]string{"ShowKmsKeyResponse", string(data)}, " ")
}

type ShowKmsKeyResponseStatus struct {
	value string
}

type ShowKmsKeyResponseStatusEnum struct {
	ENABLED          ShowKmsKeyResponseStatus
	TO_BE_ACTIVATED  ShowKmsKeyResponseStatus
	DISABLED         ShowKmsKeyResponseStatus
	PENDING_DELETION ShowKmsKeyResponseStatus
	PENDING_IMPORT   ShowKmsKeyResponseStatus
	DELETED          ShowKmsKeyResponseStatus
}

func GetShowKmsKeyResponseStatusEnum() ShowKmsKeyResponseStatusEnum {
	return ShowKmsKeyResponseStatusEnum{
		ENABLED: ShowKmsKeyResponseStatus{
			value: "ENABLED",
		},
		TO_BE_ACTIVATED: ShowKmsKeyResponseStatus{
			value: "TO_BE_ACTIVATED",
		},
		DISABLED: ShowKmsKeyResponseStatus{
			value: "DISABLED",
		},
		PENDING_DELETION: ShowKmsKeyResponseStatus{
			value: "PENDING_DELETION",
		},
		PENDING_IMPORT: ShowKmsKeyResponseStatus{
			value: "PENDING_IMPORT",
		},
		DELETED: ShowKmsKeyResponseStatus{
			value: "DELETED",
		},
	}
}

func (c ShowKmsKeyResponseStatus) Value() string {
	return c.value
}

func (c ShowKmsKeyResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowKmsKeyResponseStatus) UnmarshalJSON(b []byte) error {
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
