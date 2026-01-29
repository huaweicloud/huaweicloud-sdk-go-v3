package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowProjectVersionsCountResponse Response Object
type ShowProjectVersionsCountResponse struct {

	// **参数解释**： 查询项目下的版本数量成功或失败的状态。 **取值范围**： - success：查询项目下的版本数量成功。 - error：查询项目下的版本数量失败。
	Status *ShowProjectVersionsCountResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求的唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *ResultCount `json:"result,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowProjectVersionsCountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectVersionsCountResponse struct{}"
	}

	return strings.Join([]string{"ShowProjectVersionsCountResponse", string(data)}, " ")
}

type ShowProjectVersionsCountResponseStatus struct {
	value string
}

type ShowProjectVersionsCountResponseStatusEnum struct {
	SUCCESS ShowProjectVersionsCountResponseStatus
	ERROR   ShowProjectVersionsCountResponseStatus
}

func GetShowProjectVersionsCountResponseStatusEnum() ShowProjectVersionsCountResponseStatusEnum {
	return ShowProjectVersionsCountResponseStatusEnum{
		SUCCESS: ShowProjectVersionsCountResponseStatus{
			value: "success",
		},
		ERROR: ShowProjectVersionsCountResponseStatus{
			value: "error",
		},
	}
}

func (c ShowProjectVersionsCountResponseStatus) Value() string {
	return c.value
}

func (c ShowProjectVersionsCountResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowProjectVersionsCountResponseStatus) UnmarshalJSON(b []byte) error {
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
