package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePasswordAuthRequest Request Object
type CreatePasswordAuthRequest struct {

	// 仓库类型。 支持口令授权的仓库类型有：github、devcloud、bitbucket。
	RepoType CreatePasswordAuthRequestRepoType `json:"repo_type"`

	Body *AccessPassword `json:"body,omitempty"`
}

func (o CreatePasswordAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePasswordAuthRequest struct{}"
	}

	return strings.Join([]string{"CreatePasswordAuthRequest", string(data)}, " ")
}

type CreatePasswordAuthRequestRepoType struct {
	value string
}

type CreatePasswordAuthRequestRepoTypeEnum struct {
	GITHUB    CreatePasswordAuthRequestRepoType
	DEVCLOUD  CreatePasswordAuthRequestRepoType
	BITBUCKET CreatePasswordAuthRequestRepoType
}

func GetCreatePasswordAuthRequestRepoTypeEnum() CreatePasswordAuthRequestRepoTypeEnum {
	return CreatePasswordAuthRequestRepoTypeEnum{
		GITHUB: CreatePasswordAuthRequestRepoType{
			value: "github",
		},
		DEVCLOUD: CreatePasswordAuthRequestRepoType{
			value: "devcloud",
		},
		BITBUCKET: CreatePasswordAuthRequestRepoType{
			value: "bitbucket",
		},
	}
}

func (c CreatePasswordAuthRequestRepoType) Value() string {
	return c.value
}

func (c CreatePasswordAuthRequestRepoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePasswordAuthRequestRepoType) UnmarshalJSON(b []byte) error {
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
