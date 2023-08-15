package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PublisherSnake struct {

	// 发布者id
	PublisherId *string `json:"publisher_id,omitempty"`

	// 发布者名称
	PublisherName *string `json:"publisher_name,omitempty"`

	// 发布者展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 插件作者状态 - DISABLED 验证不通过 - VERIFIED 验证通过
	PublisherStatus *PublisherSnakePublisherStatus `json:"publisher_status,omitempty"`

	// 发布者邮箱
	Email *string `json:"email,omitempty"`

	// 网页url
	WebUrl *string `json:"web_url,omitempty"`

	// 是否开源
	Open *bool `json:"open,omitempty"`
}

func (o PublisherSnake) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublisherSnake struct{}"
	}

	return strings.Join([]string{"PublisherSnake", string(data)}, " ")
}

type PublisherSnakePublisherStatus struct {
	value string
}

type PublisherSnakePublisherStatusEnum struct {
	DISABLED PublisherSnakePublisherStatus
	VERIFIED PublisherSnakePublisherStatus
}

func GetPublisherSnakePublisherStatusEnum() PublisherSnakePublisherStatusEnum {
	return PublisherSnakePublisherStatusEnum{
		DISABLED: PublisherSnakePublisherStatus{
			value: "DISABLED",
		},
		VERIFIED: PublisherSnakePublisherStatus{
			value: "VERIFIED",
		},
	}
}

func (c PublisherSnakePublisherStatus) Value() string {
	return c.value
}

func (c PublisherSnakePublisherStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PublisherSnakePublisherStatus) UnmarshalJSON(b []byte) error {
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
