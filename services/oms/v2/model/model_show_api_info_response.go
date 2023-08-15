package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowApiInfoResponse Response Object
type ShowApiInfoResponse struct {

	// 版本号，例如v1。
	Id *string `json:"id,omitempty"`

	// 链接地址信息。
	Links *[]Link `json:"links,omitempty"`

	// 版本状态。  取值“CURRENT”，表示该版本为主推版本。  取值\"SUPPORTED\"，表示支持该版本。  取值“DEPRECATED”，表示为废弃版本，存在后续删除的可能。
	Status *ShowApiInfoResponseStatus `json:"status,omitempty"`

	// 版本更新时间。  格式为“yyyy-mm-ddThh:mm:ssZ”。  其中，T指某个时间的开始；Z指UTC时间。
	Updated        *string `json:"updated,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowApiInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowApiInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowApiInfoResponse", string(data)}, " ")
}

type ShowApiInfoResponseStatus struct {
	value string
}

type ShowApiInfoResponseStatusEnum struct {
	CURRENT    ShowApiInfoResponseStatus
	DEPRECATED ShowApiInfoResponseStatus
	SUPPORTED  ShowApiInfoResponseStatus
}

func GetShowApiInfoResponseStatusEnum() ShowApiInfoResponseStatusEnum {
	return ShowApiInfoResponseStatusEnum{
		CURRENT: ShowApiInfoResponseStatus{
			value: "CURRENT",
		},
		DEPRECATED: ShowApiInfoResponseStatus{
			value: "DEPRECATED",
		},
		SUPPORTED: ShowApiInfoResponseStatus{
			value: "SUPPORTED",
		},
	}
}

func (c ShowApiInfoResponseStatus) Value() string {
	return c.value
}

func (c ShowApiInfoResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowApiInfoResponseStatus) UnmarshalJSON(b []byte) error {
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
