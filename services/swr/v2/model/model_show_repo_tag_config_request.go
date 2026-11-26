package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepoTagConfigRequest Request Object
type ShowRepoTagConfigRequest struct {

	// **参数解释**： 消息体的类型（格式），用于指定请求体的MIME类型。 **约束限制**： 不涉及。 **取值范围**： - application/json;charset=utf-8：使用UTF-8字符集的JSON格式 - application/json：标准JSON格式 **默认取值**： 不涉及。
	ContentType ShowRepoTagConfigRequestContentType `json:"Content-Type"`

	// **参数解释**： 组织名称。用于唯一标识一个组织，组织下的所有镜像仓库都属于该组织。  **约束限制**： 不涉及。  **取值范围**： 小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，不同类型分隔符不能直接相连），小写字母或数字结尾，长度1~64个字符。  **默认取值**： 不涉及。
	Namespace string `json:"namespace"`

	// **参数解释**： 镜像仓库名称过滤关键字，用于模糊匹配镜像仓库列表。  **约束限制**： 不涉及。  **取值范围**： 小写字母或数字开头，后面跟小写字母、数字、小数点、斜杠、下划线或中划线（其中下划线最多允许连续两个，不同类型分隔符不能直接相连），小写字母或数字结尾，长度1~128个字符。  **默认取值**： 不涉及。
	Repository string `json:"repository"`

	// **参数解释**： 镜像版本名称。用于唯一标识镜像仓库下的一个镜像版本。  **约束限制**： 不涉及。  **取值范围**： 小写字母或数字开头，后面跟小写字母、数字、小数点、斜杠、下划线或中划线（其中下划线最多允许连续两个，不同类型分隔符不能直接相连），小写字母或数字结尾，长度1~128个字符。  **默认取值**： 不涉及。
	Tag string `json:"tag"`
}

func (o ShowRepoTagConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoTagConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowRepoTagConfigRequest", string(data)}, " ")
}

type ShowRepoTagConfigRequestContentType struct {
	value string
}

type ShowRepoTagConfigRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 ShowRepoTagConfigRequestContentType
	APPLICATION_JSON             ShowRepoTagConfigRequestContentType
}

func GetShowRepoTagConfigRequestContentTypeEnum() ShowRepoTagConfigRequestContentTypeEnum {
	return ShowRepoTagConfigRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: ShowRepoTagConfigRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: ShowRepoTagConfigRequestContentType{
			value: "application/json",
		},
	}
}

func (c ShowRepoTagConfigRequestContentType) Value() string {
	return c.value
}

func (c ShowRepoTagConfigRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepoTagConfigRequestContentType) UnmarshalJSON(b []byte) error {
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
