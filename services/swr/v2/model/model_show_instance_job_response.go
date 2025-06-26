package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceJobResponse Response Object
type ShowInstanceJobResponse struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 相关实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 状态
	Status *ShowInstanceJobResponseStatus `json:"status,omitempty"`

	// 失败原因
	Reason *string `json:"reason,omitempty"`

	// 实例创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 实例更新时间
	UpdatedAt      *string `json:"updated_at,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceJobResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceJobResponse", string(data)}, " ")
}

type ShowInstanceJobResponseStatus struct {
	value string
}

type ShowInstanceJobResponseStatusEnum struct {
	CREATING     ShowInstanceJobResponseStatus
	INITIALIZING ShowInstanceJobResponseStatus
	RUNNING      ShowInstanceJobResponseStatus
	FAILED       ShowInstanceJobResponseStatus
	SUCCESS      ShowInstanceJobResponseStatus
}

func GetShowInstanceJobResponseStatusEnum() ShowInstanceJobResponseStatusEnum {
	return ShowInstanceJobResponseStatusEnum{
		CREATING: ShowInstanceJobResponseStatus{
			value: "Creating",
		},
		INITIALIZING: ShowInstanceJobResponseStatus{
			value: "Initializing",
		},
		RUNNING: ShowInstanceJobResponseStatus{
			value: "Running",
		},
		FAILED: ShowInstanceJobResponseStatus{
			value: "Failed",
		},
		SUCCESS: ShowInstanceJobResponseStatus{
			value: "Success",
		},
	}
}

func (c ShowInstanceJobResponseStatus) Value() string {
	return c.value
}

func (c ShowInstanceJobResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceJobResponseStatus) UnmarshalJSON(b []byte) error {
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
