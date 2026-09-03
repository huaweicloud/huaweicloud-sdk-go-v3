package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AttachDesktopPoolUserResponse Response Object
type AttachDesktopPoolUserResponse struct {

	// CREATING：桌面创建中；WAITING：动态池排队等待；EXCEEDED：静态池已达最大值；ASSIGNING：有空闲桌面，分配中；RESETTING 重置中。
	Type *AttachDesktopPoolUserResponseType `json:"type,omitempty"`

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o AttachDesktopPoolUserResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDesktopPoolUserResponse struct{}"
	}

	return strings.Join([]string{"AttachDesktopPoolUserResponse", string(data)}, " ")
}

type AttachDesktopPoolUserResponseType struct {
	value string
}

type AttachDesktopPoolUserResponseTypeEnum struct {
	CREATING  AttachDesktopPoolUserResponseType
	WAITING   AttachDesktopPoolUserResponseType
	EXCEEDED  AttachDesktopPoolUserResponseType
	ASSIGNING AttachDesktopPoolUserResponseType
	RESETTING AttachDesktopPoolUserResponseType
}

func GetAttachDesktopPoolUserResponseTypeEnum() AttachDesktopPoolUserResponseTypeEnum {
	return AttachDesktopPoolUserResponseTypeEnum{
		CREATING: AttachDesktopPoolUserResponseType{
			value: "CREATING",
		},
		WAITING: AttachDesktopPoolUserResponseType{
			value: "WAITING",
		},
		EXCEEDED: AttachDesktopPoolUserResponseType{
			value: "EXCEEDED",
		},
		ASSIGNING: AttachDesktopPoolUserResponseType{
			value: "ASSIGNING",
		},
		RESETTING: AttachDesktopPoolUserResponseType{
			value: "RESETTING",
		},
	}
}

func (c AttachDesktopPoolUserResponseType) Value() string {
	return c.value
}

func (c AttachDesktopPoolUserResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AttachDesktopPoolUserResponseType) UnmarshalJSON(b []byte) error {
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
