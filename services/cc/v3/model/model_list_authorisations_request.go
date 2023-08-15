package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListAuthorisationsRequest Request Object
type ListAuthorisationsRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 根据ID过滤授权列表。
	Id *[]string `json:"id,omitempty"`

	// 根据名称过滤授权列表。
	Name *[]string `json:"name,omitempty"`

	// 根据描述过滤授权列表。
	Description *[]string `json:"description,omitempty"`

	// 根据云连接实例ID过滤授权列表。
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty"`

	// 根据实例ID过滤授权列表。
	InstanceId *[]ListAuthorisationsRequestInstanceId `json:"instance_id,omitempty"`
}

func (o ListAuthorisationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuthorisationsRequest struct{}"
	}

	return strings.Join([]string{"ListAuthorisationsRequest", string(data)}, " ")
}

type ListAuthorisationsRequestInstanceId struct {
	value string
}

type ListAuthorisationsRequestInstanceIdEnum struct {
	ACTIVE ListAuthorisationsRequestInstanceId
}

func GetListAuthorisationsRequestInstanceIdEnum() ListAuthorisationsRequestInstanceIdEnum {
	return ListAuthorisationsRequestInstanceIdEnum{
		ACTIVE: ListAuthorisationsRequestInstanceId{
			value: "ACTIVE",
		},
	}
}

func (c ListAuthorisationsRequestInstanceId) Value() string {
	return c.value
}

func (c ListAuthorisationsRequestInstanceId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListAuthorisationsRequestInstanceId) UnmarshalJSON(b []byte) error {
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
