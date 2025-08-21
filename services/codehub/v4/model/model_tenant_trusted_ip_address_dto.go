package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TenantTrustedIpAddressDto **参数解释：** ip白名单信息。
type TenantTrustedIpAddressDto struct {

	// **参数解释：** ip白名单id。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 用户id。
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释：** 租户id。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** ip范围。
	IpRange *string `json:"ip_range,omitempty"`

	// **参数解释：** 格式类型。 - 0，表示指定IP。 - 1，表示ip范围。 - 2，表示CIDR。
	IpType *TenantTrustedIpAddressDtoIpType `json:"ip_type,omitempty"`

	// **参数解释：** 起始ip。
	IpStart *string `json:"ip_start,omitempty"`

	// **参数解释：** 结束ip。
	IpEnd *string `json:"ip_end,omitempty"`

	// **参数解释：** 是否允许访问代码仓库。 - 0，表示禁止访问。 - 1，表示允许访问。
	ViewFlag *TenantTrustedIpAddressDtoViewFlag `json:"view_flag,omitempty"`

	// **参数解释：** 是否允许下载代码。 - 0，表示禁止下载。 - 1，表示允许下载。
	DownloadFlag *TenantTrustedIpAddressDtoDownloadFlag `json:"download_flag,omitempty"`

	// **参数解释：** 是否允许提交代码。 - 0，表示禁止提交。 - 1，表示允许提交。
	UploadFlag *TenantTrustedIpAddressDtoUploadFlag `json:"upload_flag,omitempty"`

	// **参数解释：** 备注。 **取值范围：** 字符串长度不少于0，不超过200。
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 创建时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 排序。 - 0，表示默认规则。 - 1，表示自定义规则。
	OrderFlag *TenantTrustedIpAddressDtoOrderFlag `json:"order_flag,omitempty"`
}

func (o TenantTrustedIpAddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TenantTrustedIpAddressDto struct{}"
	}

	return strings.Join([]string{"TenantTrustedIpAddressDto", string(data)}, " ")
}

type TenantTrustedIpAddressDtoIpType struct {
	value int32
}

type TenantTrustedIpAddressDtoIpTypeEnum struct {
	E_0 TenantTrustedIpAddressDtoIpType
	E_1 TenantTrustedIpAddressDtoIpType
	E_2 TenantTrustedIpAddressDtoIpType
}

func GetTenantTrustedIpAddressDtoIpTypeEnum() TenantTrustedIpAddressDtoIpTypeEnum {
	return TenantTrustedIpAddressDtoIpTypeEnum{
		E_0: TenantTrustedIpAddressDtoIpType{
			value: 0,
		}, E_1: TenantTrustedIpAddressDtoIpType{
			value: 1,
		}, E_2: TenantTrustedIpAddressDtoIpType{
			value: 2,
		},
	}
}

func (c TenantTrustedIpAddressDtoIpType) Value() int32 {
	return c.value
}

func (c TenantTrustedIpAddressDtoIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TenantTrustedIpAddressDtoIpType) UnmarshalJSON(b []byte) error {
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

type TenantTrustedIpAddressDtoViewFlag struct {
	value int32
}

type TenantTrustedIpAddressDtoViewFlagEnum struct {
	E_0 TenantTrustedIpAddressDtoViewFlag
	E_1 TenantTrustedIpAddressDtoViewFlag
}

func GetTenantTrustedIpAddressDtoViewFlagEnum() TenantTrustedIpAddressDtoViewFlagEnum {
	return TenantTrustedIpAddressDtoViewFlagEnum{
		E_0: TenantTrustedIpAddressDtoViewFlag{
			value: 0,
		}, E_1: TenantTrustedIpAddressDtoViewFlag{
			value: 1,
		},
	}
}

func (c TenantTrustedIpAddressDtoViewFlag) Value() int32 {
	return c.value
}

func (c TenantTrustedIpAddressDtoViewFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TenantTrustedIpAddressDtoViewFlag) UnmarshalJSON(b []byte) error {
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

type TenantTrustedIpAddressDtoDownloadFlag struct {
	value int32
}

type TenantTrustedIpAddressDtoDownloadFlagEnum struct {
	E_0 TenantTrustedIpAddressDtoDownloadFlag
	E_1 TenantTrustedIpAddressDtoDownloadFlag
}

func GetTenantTrustedIpAddressDtoDownloadFlagEnum() TenantTrustedIpAddressDtoDownloadFlagEnum {
	return TenantTrustedIpAddressDtoDownloadFlagEnum{
		E_0: TenantTrustedIpAddressDtoDownloadFlag{
			value: 0,
		}, E_1: TenantTrustedIpAddressDtoDownloadFlag{
			value: 1,
		},
	}
}

func (c TenantTrustedIpAddressDtoDownloadFlag) Value() int32 {
	return c.value
}

func (c TenantTrustedIpAddressDtoDownloadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TenantTrustedIpAddressDtoDownloadFlag) UnmarshalJSON(b []byte) error {
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

type TenantTrustedIpAddressDtoUploadFlag struct {
	value int32
}

type TenantTrustedIpAddressDtoUploadFlagEnum struct {
	E_0 TenantTrustedIpAddressDtoUploadFlag
	E_1 TenantTrustedIpAddressDtoUploadFlag
}

func GetTenantTrustedIpAddressDtoUploadFlagEnum() TenantTrustedIpAddressDtoUploadFlagEnum {
	return TenantTrustedIpAddressDtoUploadFlagEnum{
		E_0: TenantTrustedIpAddressDtoUploadFlag{
			value: 0,
		}, E_1: TenantTrustedIpAddressDtoUploadFlag{
			value: 1,
		},
	}
}

func (c TenantTrustedIpAddressDtoUploadFlag) Value() int32 {
	return c.value
}

func (c TenantTrustedIpAddressDtoUploadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TenantTrustedIpAddressDtoUploadFlag) UnmarshalJSON(b []byte) error {
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

type TenantTrustedIpAddressDtoOrderFlag struct {
	value int32
}

type TenantTrustedIpAddressDtoOrderFlagEnum struct {
	E_0 TenantTrustedIpAddressDtoOrderFlag
	E_1 TenantTrustedIpAddressDtoOrderFlag
}

func GetTenantTrustedIpAddressDtoOrderFlagEnum() TenantTrustedIpAddressDtoOrderFlagEnum {
	return TenantTrustedIpAddressDtoOrderFlagEnum{
		E_0: TenantTrustedIpAddressDtoOrderFlag{
			value: 0,
		}, E_1: TenantTrustedIpAddressDtoOrderFlag{
			value: 1,
		},
	}
}

func (c TenantTrustedIpAddressDtoOrderFlag) Value() int32 {
	return c.value
}

func (c TenantTrustedIpAddressDtoOrderFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TenantTrustedIpAddressDtoOrderFlag) UnmarshalJSON(b []byte) error {
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
