package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SourceRepoType 代码仓类型，支持GitHub、GitLab、Gitee、Bitbucket。
type SourceRepoType struct {
	value string
}

type SourceRepoTypeEnum struct {
	GIT_HUB   SourceRepoType
	GIT_LAB   SourceRepoType
	GITEE     SourceRepoType
	BITBUCKET SourceRepoType
}

func GetSourceRepoTypeEnum() SourceRepoTypeEnum {
	return SourceRepoTypeEnum{
		GIT_HUB: SourceRepoType{
			value: "GitHub",
		},
		GIT_LAB: SourceRepoType{
			value: "GitLab",
		},
		GITEE: SourceRepoType{
			value: "Gitee",
		},
		BITBUCKET: SourceRepoType{
			value: "Bitbucket",
		},
	}
}

func (c SourceRepoType) Value() string {
	return c.value
}

func (c SourceRepoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceRepoType) UnmarshalJSON(b []byte) error {
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
