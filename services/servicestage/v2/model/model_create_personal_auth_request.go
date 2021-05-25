package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreatePersonalAuthRequest struct {
	// 仓库类型。 支持私人令牌授权的仓库类型有：github、gitlab、gitee。

	RepoType CreatePersonalAuthRequestRepoType `json:"repo_type"`

	Body *AccessToken `json:"body,omitempty"`
}

func (o CreatePersonalAuthRequest) String() string {
	data, err := json.Marshal(o)
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

func (c CreatePersonalAuthRequestRepoType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *CreatePersonalAuthRequestRepoType) UnmarshalJSON(b []byte) error {
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
