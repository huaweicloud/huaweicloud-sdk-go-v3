package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowProjectListResponse Response Object
type ShowProjectListResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowProjectListResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// **参数解释**： 仓库列表。 **取值范围**： 不涉及。
	Result         *[]ShowProjectListResult `json:"result,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowProjectListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectListResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectListResponse", string(data)}, " ")
}

type ShowProjectListResponseStatus struct {
	value string
}

type ShowProjectListResponseStatusEnum struct {
	SUCCESS ShowProjectListResponseStatus
	ERROR   ShowProjectListResponseStatus
}

func GetShowProjectListResponseStatusEnum() ShowProjectListResponseStatusEnum {
	return ShowProjectListResponseStatusEnum{
		SUCCESS: ShowProjectListResponseStatus{
			value: "success",
		},
		ERROR: ShowProjectListResponseStatus{
			value: "error",
		},
	}
}

func (c ShowProjectListResponseStatus) Value() string {
	return c.value
}

func (c ShowProjectListResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectListResponseStatus) UnmarshalJSON(b []byte) error {
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
