package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateRepoTagRequest Request Object
type CreateRepoTagRequest struct {

	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json
	ContentType CreateRepoTagRequestContentType `json:"Content-Type"`

	// 组织名称。小写字母开头，后面跟小写字母、数字、小数点、下划线或中划线（其中下划线最多允许连续两个，小数点、下划线、中划线不能直接相连），小写字母或数字结尾，1-64个字符。
	Namespace string `json:"namespace"`

	// 镜像仓库名称
	Repository string `json:"repository"`

	Body *CreateRepoTagRequestBody `json:"body,omitempty"`
}

func (o CreateRepoTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRepoTagRequest struct{}"
	}

	return strings.Join([]string{"CreateRepoTagRequest", string(data)}, " ")
}

type CreateRepoTagRequestContentType struct {
	value string
}

type CreateRepoTagRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateRepoTagRequestContentType
	APPLICATION_JSON             CreateRepoTagRequestContentType
}

func GetCreateRepoTagRequestContentTypeEnum() CreateRepoTagRequestContentTypeEnum {
	return CreateRepoTagRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateRepoTagRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateRepoTagRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateRepoTagRequestContentType) Value() string {
	return c.value
}

func (c CreateRepoTagRequestContentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRepoTagRequestContentType) UnmarshalJSON(b []byte) error {
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
