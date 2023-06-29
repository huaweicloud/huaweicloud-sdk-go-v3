package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ListBandwidthPackagesRequest Request Object
type ListBandwidthPackagesRequest struct {

	// 分页查询时，每页返回的个数。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询时，上一页最后一条记录的ID，为空时为查询第一页。 使用说明：必须与limit一起使用。
	Marker *string `json:"marker,omitempty"`

	// 根据ID过滤带宽包实例列表。
	Id *[]string `json:"id,omitempty"`

	// 根据名称过滤带宽包实例列表。
	Name *[]string `json:"name,omitempty"`

	// 根据状态过滤带宽包实例列表。ACTIVE：表示状态可用。
	Status *[]ListBandwidthPackagesRequestStatus `json:"status,omitempty"`

	// 根据企业项目ID过滤带宽包实例列表。
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

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
