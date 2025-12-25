package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateArtifactoryResponse Response Object
type CreateArtifactoryResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *CreateArtifactoryResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *CreateArtifactoryResult `json:"result,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o CreateArtifactoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArtifactoryResponse struct{}"
	}

	return strings.Join([]string{"CreateArtifactoryResponse", string(data)}, " ")
}

type CreateArtifactoryResponseStatus struct {
	value string
}

type CreateArtifactoryResponseStatusEnum struct {
	SUCCESS CreateArtifactoryResponseStatus
	ERROR   CreateArtifactoryResponseStatus
}

func GetCreateArtifactoryResponseStatusEnum() CreateArtifactoryResponseStatusEnum {
	return CreateArtifactoryResponseStatusEnum{
		SUCCESS: CreateArtifactoryResponseStatus{
			value: "success",
		},
		ERROR: CreateArtifactoryResponseStatus{
			value: "error",
		},
	}
}

func (c CreateArtifactoryResponseStatus) Value() string {
	return c.value
}

func (c CreateArtifactoryResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateArtifactoryResponseStatus) UnmarshalJSON(b []byte) error {
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
