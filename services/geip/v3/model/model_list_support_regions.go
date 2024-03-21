package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ListSupportRegions struct {

	// 域弹性公网IP支持绑定的Region限制的ID
	Id *string `json:"id,omitempty"`

	// 支持绑定的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// 接入点信息
	AccessSite *string `json:"access_site,omitempty"`

	// region_id
	RegionId *string `json:"region_id,omitempty"`

	// - 功能说明：表示中心站点资源或者边缘站点资源 - 取值范围：center、边缘站点名称
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// remote_endpoint
	RemoteEndpoint *string `json:"remote_endpoint,omitempty"`

	// 状态
	Status *ListSupportRegionsStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ListSupportRegions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportRegions struct{}"
	}

	return strings.Join([]string{"ListSupportRegions", string(data)}, " ")
}

type ListSupportRegionsStatus struct {
	value string
}

type ListSupportRegionsStatusEnum struct {
	ACTIVE   ListSupportRegionsStatus
	INACTIVE ListSupportRegionsStatus
}

func GetListSupportRegionsStatusEnum() ListSupportRegionsStatusEnum {
	return ListSupportRegionsStatusEnum{
		ACTIVE: ListSupportRegionsStatus{
			value: "ACTIVE",
		},
		INACTIVE: ListSupportRegionsStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListSupportRegionsStatus) Value() string {
	return c.value
}

func (c ListSupportRegionsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportRegionsStatus) UnmarshalJSON(b []byte) error {
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
