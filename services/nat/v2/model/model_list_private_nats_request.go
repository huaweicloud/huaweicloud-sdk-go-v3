package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListPrivateNatsRequest Request Object
type ListPrivateNatsRequest struct {

	// 功能说明：每页返回的个数。 取值范围：0~2000。 默认值：2000。
	Limit *int32 `json:"limit,omitempty"`

	// 功能说明：分页查询起始的资源ID，为空时查询第一页。 值从上一次查询的PageInfo中的next_marker或者previous_marker中获取。
	Marker *string `json:"marker,omitempty"`

	// 是否查询前一页。
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 私网NAT网关实例的ID。
	Id *[]string `json:"id,omitempty"`

	// 私网NAT网关实例的名字。
	Name *[]string `json:"name,omitempty"`

	// 私网NAT网关实例的描述。
	Description *[]string `json:"description,omitempty"`

	// 私网NAT网关实例的规格。 取值为： \"Small\"：小型 \"Medium\"：中型 \"Large\"：大型 \"Extra-large\"：超大型
	Spec *[]ListPrivateNatsRequestSpec `json:"spec,omitempty"`

	// 私网NAT网关实例的状态。 取值为： \"ACTIVE\"：正常运行 \"FROZEN\"：冻结
	Status *[]ListPrivateNatsRequestStatus `json:"status,omitempty"`

	// 私网NAT网关实例所属VPC的ID。
	VpcId *[]string `json:"vpc_id,omitempty"`

	// 私网NAT网关实例所属子网的ID。
	VirsubnetId *[]string `json:"virsubnet_id,omitempty"`

	// 企业项目ID。创建私网NAT网关实例时，关联的企业项目ID。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
}

func (o ListPrivateNatsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPrivateNatsRequest struct{}"
	}

	return strings.Join([]string{"ListPrivateNatsRequest", string(data)}, " ")
}

type ListPrivateNatsRequestSpec struct {
	value string
}

type ListPrivateNatsRequestSpecEnum struct {
	SMALL       ListPrivateNatsRequestSpec
	MEDIUM      ListPrivateNatsRequestSpec
	LARGE       ListPrivateNatsRequestSpec
	EXTRA_LARGE ListPrivateNatsRequestSpec
}

func GetListPrivateNatsRequestSpecEnum() ListPrivateNatsRequestSpecEnum {
	return ListPrivateNatsRequestSpecEnum{
		SMALL: ListPrivateNatsRequestSpec{
			value: "Small",
		},
		MEDIUM: ListPrivateNatsRequestSpec{
			value: "Medium",
		},
		LARGE: ListPrivateNatsRequestSpec{
			value: "Large",
		},
		EXTRA_LARGE: ListPrivateNatsRequestSpec{
			value: "Extra-large",
		},
	}
}

func (c ListPrivateNatsRequestSpec) Value() string {
	return c.value
}

func (c ListPrivateNatsRequestSpec) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateNatsRequestSpec) UnmarshalJSON(b []byte) error {
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

type ListPrivateNatsRequestStatus struct {
	value string
}

type ListPrivateNatsRequestStatusEnum struct {
	ACTIVE ListPrivateNatsRequestStatus
	FROZEN ListPrivateNatsRequestStatus
}

func GetListPrivateNatsRequestStatusEnum() ListPrivateNatsRequestStatusEnum {
	return ListPrivateNatsRequestStatusEnum{
		ACTIVE: ListPrivateNatsRequestStatus{
			value: "ACTIVE",
		},
		FROZEN: ListPrivateNatsRequestStatus{
			value: "FROZEN",
		},
	}
}

func (c ListPrivateNatsRequestStatus) Value() string {
	return c.value
}

func (c ListPrivateNatsRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListPrivateNatsRequestStatus) UnmarshalJSON(b []byte) error {
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
