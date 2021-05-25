package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateUserRepositoryAuthRequest struct {
	// 消息体的类型（格式），下方类型可任选其一使用： application/json;charset=utf-8 application/json

	ContentType CreateUserRepositoryAuthRequestContentType `json:"Content-Type"`
	// 组织名称

	Namespace string `json:"namespace"`
	// 镜像仓库名称

	Repository string `json:"repository"`

	Body *[]UserAuth `json:"body,omitempty"`
}

func (o CreateUserRepositoryAuthRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateUserRepositoryAuthRequest struct{}"
	}

	return strings.Join([]string{"CreateUserRepositoryAuthRequest", string(data)}, " ")
}

type CreateUserRepositoryAuthRequestContentType struct {
	value string
}

type CreateUserRepositoryAuthRequestContentTypeEnum struct {
	APPLICATION_JSONCHARSETUTF_8 CreateUserRepositoryAuthRequestContentType
	APPLICATION_JSON             CreateUserRepositoryAuthRequestContentType
}

func GetCreateUserRepositoryAuthRequestContentTypeEnum() CreateUserRepositoryAuthRequestContentTypeEnum {
	return CreateUserRepositoryAuthRequestContentTypeEnum{
		APPLICATION_JSONCHARSETUTF_8: CreateUserRepositoryAuthRequestContentType{
			value: "application/json;charset=utf-8",
		},
		APPLICATION_JSON: CreateUserRepositoryAuthRequestContentType{
			value: "application/json",
		},
	}
}

func (c CreateUserRepositoryAuthRequestContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreateUserRepositoryAuthRequestContentType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
