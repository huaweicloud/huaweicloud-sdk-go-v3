package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePersonalAuthRequest Request Object
type CreatePersonalAuthRequest struct {

	// 仓库类型。 支持私人令牌授权的仓库类型有：github、gitlab、gitee。
	RepoType CreatePersonalAuthRequestRepoType `json:"repo_type"`

	Body *AccessToken `json:"body,omitempty"`
}

func (o CreatePersonalAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePersonalAuthRequest struct{}"
	}

	return strings.Join([]string{"CreatePersonalAuthRequest", string(data)}, " ")
}

type CreatePersonalAuthRequestRepoType struct {
	value string
}

type CreatePersonalAuthRequestRepoTypeEnum struct {
	GITHUB CreatePersonalAuthRequestRepoType
	GITLAB CreatePersonalAuthRequestRepoType
	GITEE  CreatePersonalAuthRequestRepoType
}

func GetCreatePersonalAuthRequestRepoTypeEnum() CreatePersonalAuthRequestRepoTypeEnum {
	return CreatePersonalAuthRequestRepoTypeEnum{
		GITHUB: CreatePersonalAuthRequestRepoType{
			value: "github",
		},
		GITLAB: CreatePersonalAuthRequestRepoType{
			value: "gitlab",
		},
		GITEE: CreatePersonalAuthRequestRepoType{
			value: "gitee",
		},
	}
}

func (c CreatePersonalAuthRequestRepoType) Value() string {
	return c.value
}

func (c CreatePersonalAuthRequestRepoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePersonalAuthRequestRepoType) UnmarshalJSON(b []byte) error {
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
