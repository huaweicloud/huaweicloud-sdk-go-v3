package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ListSupportInstancesSupportRegions struct {

	// 域弹性公网IP支持绑定的Region限制的ID
	Id *string `json:"id,omitempty"`

	// 支持绑定的实例类型
	InstanceType *string `json:"instance_type,omitempty"`

	// - 功能说明：表示中心站点资源或者边缘站点资源 - 取值范围：center、边缘站点名称
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// region_id
	RegionId *string `json:"region_id,omitempty"`

	// access_site,后端实例所在的站点信息
	AccessSite *string `json:"access_site,omitempty"`

	// 状态
	Status *ListSupportInstancesSupportRegionsStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ListSupportInstancesSupportRegions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportInstancesSupportRegions struct{}"
	}

	return strings.Join([]string{"ListSupportInstancesSupportRegions", string(data)}, " ")
}

type ListSupportInstancesSupportRegionsStatus struct {
	value string
}

type ListSupportInstancesSupportRegionsStatusEnum struct {
	ACTIVE   ListSupportInstancesSupportRegionsStatus
	INACTIVE ListSupportInstancesSupportRegionsStatus
}

func GetListSupportInstancesSupportRegionsStatusEnum() ListSupportInstancesSupportRegionsStatusEnum {
	return ListSupportInstancesSupportRegionsStatusEnum{
		ACTIVE: ListSupportInstancesSupportRegionsStatus{
			value: "ACTIVE",
		},
		INACTIVE: ListSupportInstancesSupportRegionsStatus{
			value: "INACTIVE",
		},
	}
}

func (c ListSupportInstancesSupportRegionsStatus) Value() string {
	return c.value
}

func (c ListSupportInstancesSupportRegionsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportInstancesSupportRegionsStatus) UnmarshalJSON(b []byte) error {
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
