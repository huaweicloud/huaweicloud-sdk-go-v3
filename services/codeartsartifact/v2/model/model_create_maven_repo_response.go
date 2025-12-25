package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateMavenRepoResponse Response Object
type CreateMavenRepoResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *CreateMavenRepoResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *CreateMavenRepoResult `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o CreateMavenRepoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMavenRepoResponse struct{}"
	}

	return strings.Join([]string{"CreateMavenRepoResponse", string(data)}, " ")
}

type CreateMavenRepoResponseStatus struct {
	value string
}

type CreateMavenRepoResponseStatusEnum struct {
	SUCCESS CreateMavenRepoResponseStatus
	ERROR   CreateMavenRepoResponseStatus
}

func GetCreateMavenRepoResponseStatusEnum() CreateMavenRepoResponseStatusEnum {
	return CreateMavenRepoResponseStatusEnum{
		SUCCESS: CreateMavenRepoResponseStatus{
			value: "success",
		},
		ERROR: CreateMavenRepoResponseStatus{
			value: "error",
		},
	}
}

func (c CreateMavenRepoResponseStatus) Value() string {
	return c.value
}

func (c CreateMavenRepoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateMavenRepoResponseStatus) UnmarshalJSON(b []byte) error {
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
