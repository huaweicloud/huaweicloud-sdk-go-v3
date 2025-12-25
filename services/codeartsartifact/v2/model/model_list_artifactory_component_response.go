package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListArtifactoryComponentResponse Response Object
type ListArtifactoryComponentResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ListArtifactoryComponentResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**:  文件详情。 **取值范围**: 不涉及。
	Result         *interface{} `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListArtifactoryComponentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArtifactoryComponentResponse struct{}"
	}

	return strings.Join([]string{"ListArtifactoryComponentResponse", string(data)}, " ")
}

type ListArtifactoryComponentResponseStatus struct {
	value string
}

type ListArtifactoryComponentResponseStatusEnum struct {
	SUCCESS ListArtifactoryComponentResponseStatus
	ERROR   ListArtifactoryComponentResponseStatus
}

func GetListArtifactoryComponentResponseStatusEnum() ListArtifactoryComponentResponseStatusEnum {
	return ListArtifactoryComponentResponseStatusEnum{
		SUCCESS: ListArtifactoryComponentResponseStatus{
			value: "success",
		},
		ERROR: ListArtifactoryComponentResponseStatus{
			value: "error",
		},
	}
}

func (c ListArtifactoryComponentResponseStatus) Value() string {
	return c.value
}

func (c ListArtifactoryComponentResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListArtifactoryComponentResponseStatus) UnmarshalJSON(b []byte) error {
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
