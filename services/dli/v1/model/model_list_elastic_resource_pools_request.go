package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListElasticResourcePoolsRequest Request Object
type ListElasticResourcePoolsRequest struct {

	// 默认为100
	Limit *int32 `json:"limit,omitempty"`

	// 通过弹性资源池名称进行模糊匹配
	Name *string `json:"name,omitempty"`

	// 默认为0
	Offset *int32 `json:"offset,omitempty"`

	// 弹性资源池状态
	Status *ListElasticResourcePoolsRequestStatus `json:"status,omitempty"`

	// 查询根据标签进行过滤。
	Tags *string `json:"tags,omitempty"`
}

func (o ListElasticResourcePoolsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListElasticResourcePoolsRequest struct{}"
	}

	return strings.Join([]string{"ListElasticResourcePoolsRequest", string(data)}, " ")
}

type ListElasticResourcePoolsRequestStatus struct {
	value string
}

type ListElasticResourcePoolsRequestStatusEnum struct {
	AVAILABLE ListElasticResourcePoolsRequestStatus
	SCALING   ListElasticResourcePoolsRequestStatus
	CREATING  ListElasticResourcePoolsRequestStatus
	FAILED    ListElasticResourcePoolsRequestStatus
}

func GetListElasticResourcePoolsRequestStatusEnum() ListElasticResourcePoolsRequestStatusEnum {
	return ListElasticResourcePoolsRequestStatusEnum{
		AVAILABLE: ListElasticResourcePoolsRequestStatus{
			value: "AVAILABLE",
		},
		SCALING: ListElasticResourcePoolsRequestStatus{
			value: "SCALING",
		},
		CREATING: ListElasticResourcePoolsRequestStatus{
			value: "CREATING",
		},
		FAILED: ListElasticResourcePoolsRequestStatus{
			value: "FAILED",
		},
	}
}

func (c ListElasticResourcePoolsRequestStatus) Value() string {
	return c.value
}

func (c ListElasticResourcePoolsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListElasticResourcePoolsRequestStatus) UnmarshalJSON(b []byte) error {
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
