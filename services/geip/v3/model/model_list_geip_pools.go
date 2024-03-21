package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ListGeipPools struct {

	// 全域弹性公网IP池的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：全域弹性公网IP池名称 - 取值范围：1-64，支持数字、字母、中文、_(下划线)、-（中划线）、.（点）
	Name *string `json:"name,omitempty"`

	// 英文名称
	EnName *string `json:"en_name,omitempty"`

	// 中文名称
	CnName *string `json:"cn_name,omitempty"`

	// 状态
	Status *ListGeipPoolsStatus `json:"status,omitempty"`

	// 全域弹性公网IP所属线路
	Isp *string `json:"isp,omitempty"`

	// - 功能说明：全域弹性公网IP池的版本 - 取值范围：4、6
	IpVersion *ListGeipPoolsIpVersion `json:"ip_version,omitempty"`

	// 接入点信息
	AccessSite *string `json:"access_site,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`

	// 地址池支持的全域公网带宽类型资源
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

type ListGeipPoolsIpVersion struct {
	value int32
}

type ListGeipPoolsIpVersionEnum struct {
	E_4 ListGeipPoolsIpVersion
	E_6 ListGeipPoolsIpVersion
}

func GetListGeipPoolsIpVersionEnum() ListGeipPoolsIpVersionEnum {
	return ListGeipPoolsIpVersionEnum{
		E_4: ListGeipPoolsIpVersion{
			value: 4,
		}, E_6: ListGeipPoolsIpVersion{
			value: 6,
		},
	}
}

func (c ListGeipPoolsIpVersion) Value() int32 {
	return c.value
}

func (c ListGeipPoolsIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListGeipPoolsIpVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
