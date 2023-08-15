package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Follow302StatusBody struct {

	// 加速域名id。
	DomainId *string `json:"domain_id,omitempty"`

	// follow302状态，off：关闭，on：开启。
	FollowStatus *Follow302StatusBodyFollowStatus `json:"follow_status,omitempty"`
}

func (o Follow302StatusBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Follow302StatusBody struct{}"
	}

	return strings.Join([]string{"Follow302StatusBody", string(data)}, " ")
}

type Follow302StatusBodyFollowStatus struct {
	value string
}

type Follow302StatusBodyFollowStatusEnum struct {
	OFF Follow302StatusBodyFollowStatus
	ON  Follow302StatusBodyFollowStatus
}

func GetFollow302StatusBodyFollowStatusEnum() Follow302StatusBodyFollowStatusEnum {
	return Follow302StatusBodyFollowStatusEnum{
		OFF: Follow302StatusBodyFollowStatus{
			value: "off",
		},
		ON: Follow302StatusBodyFollowStatus{
			value: "on",
		},
	}
}

func (c Follow302StatusBodyFollowStatus) Value() string {
	return c.value
}

func (c Follow302StatusBodyFollowStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Follow302StatusBodyFollowStatus) UnmarshalJSON(b []byte) error {
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
