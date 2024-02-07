package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ListGeipPools struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 英文名称
	EnName *string `json:"en_name,omitempty"`

	// 中文名称
	CnName *string `json:"cn_name,omitempty"`

	// 状态
	Status *ListGeipPoolsStatus `json:"status,omitempty"`

	// 线路
	Isp *string `json:"isp,omitempty"`

	// IPv4或IPv6
	IpVersion *int32 `json:"ip_version,omitempty"`

	// 接入点信息
	AccessSite *string `json:"access_site,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	AllowedBandwidthTypes *[]AllowedBandwidthTypes `json:"allowed_bandwidth_types,omitempty"`
}

func (o ListGeipPools) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGeipPools struct{}"
	}

	return strings.Join([]string{"ListGeipPools", string(data)}, " ")
}

type ListGeipPoolsStatus struct {
	value string
}

type ListGeipPoolsStatusEnum struct {
	ACTIVE   ListGeipPoolsStatus
	INACTIVE ListGeipPoolsStatus
	SOLDOUT  ListGeipPoolsStatus
}

func GetListGeipPoolsStatusEnum() ListGeipPoolsStatusEnum {
	return ListGeipPoolsStatusEnum{
		ACTIVE: ListGeipPoolsStatus{
			value: "ACTIVE",
		},
		INACTIVE: ListGeipPoolsStatus{
			value: "INACTIVE",
		},
		SOLDOUT: ListGeipPoolsStatus{
			value: "SOLDOUT",
		},
	}
}

func (c ListGeipPoolsStatus) Value() string {
	return c.value
}

func (c ListGeipPoolsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsStatus) UnmarshalJSON(b []byte) error {
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
