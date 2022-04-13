package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type RunQueryInstanceResponse struct {
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

	Status *RunQueryInstanceResponseStatus `json:"status,omitempty"`
	// 实例名称。

	InstanceName   *string `json:"instanceName,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o RunQueryInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunQueryInstanceResponse struct{}"
	}

	return strings.Join([]string{"RunQueryInstanceResponse", string(data)}, " ")
}

type RunQueryInstanceResponseStatus struct {
	value string
}

type RunQueryInstanceResponseStatusEnum struct {
	NORMAL          RunQueryInstanceResponseStatus
	ARREARAGE       RunQueryInstanceResponseStatus
	CREATION        RunQueryInstanceResponseStatus
	CREATION_FAILD  RunQueryInstanceResponseStatus
	DELETING        RunQueryInstanceResponseStatus
	DELETING_FAILED RunQueryInstanceResponseStatus
	ABNORMAL        RunQueryInstanceResponseStatus
}

func GetRunQueryInstanceResponseStatusEnum() RunQueryInstanceResponseStatusEnum {
	return RunQueryInstanceResponseStatusEnum{
		NORMAL: RunQueryInstanceResponseStatus{
			value: "NORMAL",
		},
		ARREARAGE: RunQueryInstanceResponseStatus{
			value: "ARREARAGE",
		},
		CREATION: RunQueryInstanceResponseStatus{
			value: "CREATION",
		},
		CREATION_FAILD: RunQueryInstanceResponseStatus{
			value: "CREATION_FAILD",
		},
		DELETING: RunQueryInstanceResponseStatus{
			value: "DELETING",
		},
		DELETING_FAILED: RunQueryInstanceResponseStatus{
			value: "DELETING_FAILED",
		},
		ABNORMAL: RunQueryInstanceResponseStatus{
			value: "ABNORMAL",
		},
	}
}

func (c RunQueryInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunQueryInstanceResponseStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
