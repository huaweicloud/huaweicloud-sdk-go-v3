package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RunCreateInstanceResponse Response Object
type RunCreateInstanceResponse struct {

	// 模型展示名或领域名称。
	Domain *string `json:"domain,omitempty"`

	// 描述。
	Desc *string `json:"desc,omitempty"`

	// 注册时间。
	RegisterDate *int64 `json:"registerDate,omitempty"`

	// 过期时间，-1表示永不过期。
	ExpiredDate *int64 `json:"expiredDate,omitempty"`

	// 规格，即实例的图片数量规格，默认为30000000（单位：张）。
	Level *int32 `json:"level,omitempty"`

	// 图片自定义标签。
	Tags *[]string `json:"tags,omitempty"`

	// 实例的状态，有以下状态信息：   - NORMAL：正常。   - ARREARAGE：欠费。   - CREATION：创建中。   - CREATION_FAILD：创建失败。   - DELETING：删除中。   - DELETING_FAILED：删除失败。   - ABNORMAL：异常。
	Status *RunCreateInstanceResponseStatus `json:"status,omitempty"`

	// 实例名称。
	InstanceName   *string `json:"instanceName,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunCreateInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunCreateInstanceResponse struct{}"
	}

	return strings.Join([]string{"RunCreateInstanceResponse", string(data)}, " ")
}

type RunCreateInstanceResponseStatus struct {
	value string
}

type RunCreateInstanceResponseStatusEnum struct {
	NORMAL          RunCreateInstanceResponseStatus
	ARREARAGE       RunCreateInstanceResponseStatus
	CREATION        RunCreateInstanceResponseStatus
	CREATION_FAILD  RunCreateInstanceResponseStatus
	DELETING        RunCreateInstanceResponseStatus
	DELETING_FAILED RunCreateInstanceResponseStatus
	ABNORMAL        RunCreateInstanceResponseStatus
}

func GetRunCreateInstanceResponseStatusEnum() RunCreateInstanceResponseStatusEnum {
	return RunCreateInstanceResponseStatusEnum{
		NORMAL: RunCreateInstanceResponseStatus{
			value: "NORMAL",
		},
		ARREARAGE: RunCreateInstanceResponseStatus{
			value: "ARREARAGE",
		},
		CREATION: RunCreateInstanceResponseStatus{
			value: "CREATION",
		},
		CREATION_FAILD: RunCreateInstanceResponseStatus{
			value: "CREATION_FAILD",
		},
		DELETING: RunCreateInstanceResponseStatus{
			value: "DELETING",
		},
		DELETING_FAILED: RunCreateInstanceResponseStatus{
			value: "DELETING_FAILED",
		},
		ABNORMAL: RunCreateInstanceResponseStatus{
			value: "ABNORMAL",
		},
	}
}

func (c RunCreateInstanceResponseStatus) Value() string {
	return c.value
}

func (c RunCreateInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunCreateInstanceResponseStatus) UnmarshalJSON(b []byte) error {
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
