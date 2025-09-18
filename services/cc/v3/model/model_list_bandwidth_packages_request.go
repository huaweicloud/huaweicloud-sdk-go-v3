package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBandwidthPackagesRequest Request Object
type ListBandwidthPackagesRequest struct {

	// 每页返回的个数。 取值范围：1~2000。
	Limit *int32 `json:"limit,omitempty"`

	// 翻页信息，从上次API调用返回的翻页数据中获取，可填写前一页marker或者后一页marker，填入前一页previous_marker就向前翻页，后一页next_marker就向后翻页。 翻页过程中，查询条件不能修改，包括过滤条件、排序条件、limit。
	Marker *string `json:"marker,omitempty"`

	// 根据ID查询，可查询多个ID。
	Id *[]string `json:"id,omitempty"`

	// 根据名称查询，可查询多个名称。
	Name *[]string `json:"name,omitempty"`

	// 根据企业项目ID过滤列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// 根据云连接的ID过滤列表。
	CloudConnectionId *[]string `json:"cloud_connection_id,omitempty"`

	// 根据状态过滤带宽包实例列表。ACTIVE：表示状态可用。
	Status *[]ListBandwidthPackagesRequestStatus `json:"status,omitempty"`

	// 根据计费方式过滤带宽包实例列表。
	BillingMode *[]string `json:"billing_mode,omitempty"`

	// 根据绑定的资源ID过滤带宽包实例列表。
	ResourceId *[]string `json:"resource_id,omitempty"`
}

func (o ListBandwidthPackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBandwidthPackagesRequest struct{}"
	}

	return strings.Join([]string{"ListBandwidthPackagesRequest", string(data)}, " ")
}

type ListBandwidthPackagesRequestStatus struct {
	value string
}

type ListBandwidthPackagesRequestStatusEnum struct {
	ACTIVE ListBandwidthPackagesRequestStatus
}

func GetListBandwidthPackagesRequestStatusEnum() ListBandwidthPackagesRequestStatusEnum {
	return ListBandwidthPackagesRequestStatusEnum{
		ACTIVE: ListBandwidthPackagesRequestStatus{
			value: "ACTIVE",
		},
	}
}

func (c ListBandwidthPackagesRequestStatus) Value() string {
	return c.value
}

func (c ListBandwidthPackagesRequestStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListBandwidthPackagesRequestStatus) UnmarshalJSON(b []byte) error {
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
