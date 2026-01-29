package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowProjectRelatedRepositoryResponse Response Object
type ShowProjectRelatedRepositoryResponse struct {

	// **参数解释**： 查询项目列表成功或失败的状态。 **取值范围**： - success：查询项目列表成功。 - error：查询项目列表失败。
	Status *ShowProjectRelatedRepositoryResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求的唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *PrivilegeProjectInfoV5 `json:"result,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowProjectRelatedRepositoryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectRelatedRepositoryResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectRelatedRepositoryResponse", string(data)}, " ")
}

type ShowProjectRelatedRepositoryResponseStatus struct {
	value string
}

type ShowProjectRelatedRepositoryResponseStatusEnum struct {
	SUCCESS ShowProjectRelatedRepositoryResponseStatus
	ERROR   ShowProjectRelatedRepositoryResponseStatus
}

func GetShowProjectRelatedRepositoryResponseStatusEnum() ShowProjectRelatedRepositoryResponseStatusEnum {
	return ShowProjectRelatedRepositoryResponseStatusEnum{
		SUCCESS: ShowProjectRelatedRepositoryResponseStatus{
			value: "success",
		},
		ERROR: ShowProjectRelatedRepositoryResponseStatus{
			value: "error",
		},
	}
}

func (c ShowProjectRelatedRepositoryResponseStatus) Value() string {
	return c.value
}

func (c ShowProjectRelatedRepositoryResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectRelatedRepositoryResponseStatus) UnmarshalJSON(b []byte) error {
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
