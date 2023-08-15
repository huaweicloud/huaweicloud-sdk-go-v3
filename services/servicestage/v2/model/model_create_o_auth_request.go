package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateOAuthRequest Request Object
type CreateOAuthRequest struct {

	// 仓库类型。 支持OAuth授权的仓库类型有：github、gitlab、gitee、bitbucket。
	RepoType CreateOAuthRequestRepoType `json:"repo_type"`

	// 站点标签。 比如国际站的，?tag=intl。 默认为空。
	Tag *string `json:"tag,omitempty"`

	Body *OAuth `json:"body,omitempty"`
}

func (o CreateOAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOAuthRequest struct{}"
	}

	return strings.Join([]string{"CreateOAuthRequest", string(data)}, " ")
}

type CreateOAuthRequestRepoType struct {
	value string
}

type CreateOAuthRequestRepoTypeEnum struct {
	GITHUB    CreateOAuthRequestRepoType
	GITLAB    CreateOAuthRequestRepoType
	GITEE     CreateOAuthRequestRepoType
	BITBUCKET CreateOAuthRequestRepoType
}

func GetCreateOAuthRequestRepoTypeEnum() CreateOAuthRequestRepoTypeEnum {
	return CreateOAuthRequestRepoTypeEnum{
		GITHUB: CreateOAuthRequestRepoType{
			value: "github",
		},
		GITLAB: CreateOAuthRequestRepoType{
			value: "gitlab",
		},
		GITEE: CreateOAuthRequestRepoType{
			value: "gitee",
		},
		BITBUCKET: CreateOAuthRequestRepoType{
			value: "bitbucket",
		},
	}
}

func (c CreateOAuthRequestRepoType) Value() string {
	return c.value
}

func (c CreateOAuthRequestRepoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateOAuthRequestRepoType) UnmarshalJSON(b []byte) error {
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
