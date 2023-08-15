package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRetentionRequest Request Object
type CreateRetentionRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType CreateRetentionRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	Body *CreateRetentionRequestBody `json:"body,omitempty"`
}

func (o CreateRetentionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRetentionRequest struct{}"
	}

	return strings.Join([]string{"CreateRetentionRequest", string(data)}, " ")
}

type CreateRetentionRequestContentType struct {
	value string
}

type CreateRetentionRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateRetentionRequestContentType
	APPLICATION_JSON             CreateRetentionRequestContentType
}

func GetCreateRetentionRequestContentTypeEnum() CreateRetentionRequestContentTypeEnum {
	return CreateRetentionRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateRetentionRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateRetentionRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateRetentionRequestContentType) Value() string {
	return c.value
}

func (c CreateRetentionRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRetentionRequestContentType) UnmarshalJSON(b []byte) error {
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
