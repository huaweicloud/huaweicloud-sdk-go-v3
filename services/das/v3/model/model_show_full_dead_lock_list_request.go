package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFullDeadLockListRequest Request Object
type ShowFullDeadLockListRequest struct {

	// 实例id
	InstanceId string `json:"instance_id"`

	// 开始时间
	StartAt int64 `json:"start_at"`

	// 结束时间
	EndAt int64 `json:"end_at"`

	// 当前页数
	PageNum *int32 `json:"page_num,omitempty"`

	// 每页返回数量，默认10
	PageSize *int32 `json:"page_size,omitempty"`

	// 语言
	XLanguage *ShowFullDeadLockListRequestXLanguage `json:"X-Language,omitempty"`
}

func (o ShowFullDeadLockListRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFullDeadLockListRequest struct{}"
	}

	return strings.Join([]string{"ShowFullDeadLockListRequest", string(data)}, " ")
}

type ShowFullDeadLockListRequestXLanguage struct {
	value string
}

type ShowFullDeadLockListRequestXLanguageEnum struct {
	ZH_CN ShowFullDeadLockListRequestXLanguage
	EN_US ShowFullDeadLockListRequestXLanguage
}

func GetShowFullDeadLockListRequestXLanguageEnum() ShowFullDeadLockListRequestXLanguageEnum {
	return ShowFullDeadLockListRequestXLanguageEnum{
		ZH_CN: ShowFullDeadLockListRequestXLanguage{
			value: "zh-cn",
		},
		EN_US: ShowFullDeadLockListRequestXLanguage{
			value: "en-us",
		},
	}
}

func (c ShowFullDeadLockListRequestXLanguage) Value() string {
	return c.value
}

func (c ShowFullDeadLockListRequestXLanguage) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFullDeadLockListRequestXLanguage) UnmarshalJSON(b []byte) error {
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
