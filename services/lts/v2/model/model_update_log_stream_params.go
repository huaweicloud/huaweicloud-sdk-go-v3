package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateLogStreamParams struct {

	// 日志存储时间（天）。   该参数仅对华东-上海一、华北-北京四、华南-广州用户开放。
	TtlInDays UpdateLogStreamParamsTtlInDays `json:"ttl_in_days"`

	// 标签字段信息
	Tags *[]TagsBody `json:"tags,omitempty"`
}

func (o UpdateLogStreamParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogStreamParams struct{}"
	}

	return strings.Join([]string{"UpdateLogStreamParams", string(data)}, " ")
}

type UpdateLogStreamParamsTtlInDays struct {
	value int32
}

type UpdateLogStreamParamsTtlInDaysEnum struct {
	E_7 UpdateLogStreamParamsTtlInDays
}

func GetUpdateLogStreamParamsTtlInDaysEnum() UpdateLogStreamParamsTtlInDaysEnum {
	return UpdateLogStreamParamsTtlInDaysEnum{
		E_7: UpdateLogStreamParamsTtlInDays{
			value: 7,
		},
	}
}

func (c UpdateLogStreamParamsTtlInDays) Value() int32 {
	return c.value
}

func (c UpdateLogStreamParamsTtlInDays) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLogStreamParamsTtlInDays) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
