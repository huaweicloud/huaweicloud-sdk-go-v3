package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowRepoUserInfoResponse Response Object
type ShowRepoUserInfoResponse struct {

	// **参数解释**： 请求成功或失败状态。 **取值范围**： - success：请求成功。 - error：请求失败。
	Status *ShowRepoUserInfoResponseStatus `json:"status,omitempty"`

	// **参数解释**： 请求ID，当前请求唯一标识。 **取值范围**： 数字及中划线（-）组成的字符串。
	TraceId *string `json:"trace_id,omitempty"`

	Result         *RepositoryUserDo `json:"result,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowRepoUserInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRepoUserInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowRepoUserInfoResponse", string(data)}, " ")
}

type ShowRepoUserInfoResponseStatus struct {
	value string
}

type ShowRepoUserInfoResponseStatusEnum struct {
	SUCCESS ShowRepoUserInfoResponseStatus
	ERROR   ShowRepoUserInfoResponseStatus
}

func GetShowRepoUserInfoResponseStatusEnum() ShowRepoUserInfoResponseStatusEnum {
	return ShowRepoUserInfoResponseStatusEnum{
		SUCCESS: ShowRepoUserInfoResponseStatus{
			value: "success",
		},
		ERROR: ShowRepoUserInfoResponseStatus{
			value: "error",
		},
	}
}

func (c ShowRepoUserInfoResponseStatus) Value() string {
	return c.value
}

func (c ShowRepoUserInfoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowRepoUserInfoResponseStatus) UnmarshalJSON(b []byte) error {
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
