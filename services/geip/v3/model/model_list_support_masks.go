package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ListSupportMasks struct {

	// 全域弹性公网IP段支持的掩码的ID
	Id *string `json:"id,omitempty"`

	// - 功能说明：全域弹性公网IP的版本 - 取值范围：4、6
	IpVersion *ListSupportMasksIpVersion `json:"ip_version,omitempty"`

	// 掩码长度
	Mask *int32 `json:"mask,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ListSupportMasks) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSupportMasks struct{}"
	}

	return strings.Join([]string{"ListSupportMasks", string(data)}, " ")
}

type ListSupportMasksIpVersion struct {
	value int32
}

type ListSupportMasksIpVersionEnum struct {
	E_4 ListSupportMasksIpVersion
	E_6 ListSupportMasksIpVersion
}

func GetListSupportMasksIpVersionEnum() ListSupportMasksIpVersionEnum {
	return ListSupportMasksIpVersionEnum{
		E_4: ListSupportMasksIpVersion{
			value: 4,
		}, E_6: ListSupportMasksIpVersion{
			value: 6,
		},
	}
}

func (c ListSupportMasksIpVersion) Value() int32 {
	return c.value
}

func (c ListSupportMasksIpVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListSupportMasksIpVersion) UnmarshalJSON(b []byte) error {
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
