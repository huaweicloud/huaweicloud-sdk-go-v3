package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowUserTicketResponse Response Object
type ShowUserTicketResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowUserTicketResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	// 参数解释: ticket信息。 取值范围: 不涉及。
	Result         *string `json:"result,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUserTicketResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUserTicketResponse struct{}"
	}

	return strings.Join([]string{"ShowUserTicketResponse", string(data)}, " ")
}

type ShowUserTicketResponseStatus struct {
	value string
}

type ShowUserTicketResponseStatusEnum struct {
	SUCCESS ShowUserTicketResponseStatus
	ERROR   ShowUserTicketResponseStatus
}

func GetShowUserTicketResponseStatusEnum() ShowUserTicketResponseStatusEnum {
	return ShowUserTicketResponseStatusEnum{
		SUCCESS: ShowUserTicketResponseStatus{
			value: "success",
		},
		ERROR: ShowUserTicketResponseStatus{
			value: "error",
		},
	}
}

func (c ShowUserTicketResponseStatus) Value() string {
	return c.value
}

func (c ShowUserTicketResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowUserTicketResponseStatus) UnmarshalJSON(b []byte) error {
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
