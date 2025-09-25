package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SwitchKmsTdeRequestBody struct {

	// **参数解释**: kms主密钥ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	KmsTdeKeyId string `json:"kms_tde_key_id"`

	// **参数解释**: GaussDB使用透明加密的KMS主密钥ID所在资源空间名称。 获取方法请参见[获取项目名称](https://support.huaweicloud.com/api-gaussdb/gaussdb_api_196.html)。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	KmsProjectName string `json:"kms_project_name"`

	// **参数解释**: 需要切换的状态：on表示开启。 **约束限制**: 不涉及。 **取值范围**: on: 开启。 **默认取值**: on
	KmsTdeStatus SwitchKmsTdeRequestBodyKmsTdeStatus `json:"kms_tde_status"`
}

func (o SwitchKmsTdeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchKmsTdeRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchKmsTdeRequestBody", string(data)}, " ")
}

type SwitchKmsTdeRequestBodyKmsTdeStatus struct {
	value string
}

type SwitchKmsTdeRequestBodyKmsTdeStatusEnum struct {
	ON SwitchKmsTdeRequestBodyKmsTdeStatus
}

func GetSwitchKmsTdeRequestBodyKmsTdeStatusEnum() SwitchKmsTdeRequestBodyKmsTdeStatusEnum {
	return SwitchKmsTdeRequestBodyKmsTdeStatusEnum{
		ON: SwitchKmsTdeRequestBodyKmsTdeStatus{
			value: "on",
		},
	}
}

func (c SwitchKmsTdeRequestBodyKmsTdeStatus) Value() string {
	return c.value
}

func (c SwitchKmsTdeRequestBodyKmsTdeStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SwitchKmsTdeRequestBodyKmsTdeStatus) UnmarshalJSON(b []byte) error {
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
