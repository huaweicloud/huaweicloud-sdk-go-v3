package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TrustedIpAddressDto **参数解释：** ip白名单信息。
type TrustedIpAddressDto struct {

	// **参数解释：** 白名单id。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** ip范围。
	IpRange *string `json:"ip_range,omitempty"`

	// **参数解释：** 格式类型。 - 0，表示指定ip。 - 1，表示ip范围。 - 2，表示CIDR。
	IpType *TrustedIpAddressDtoIpType `json:"ip_type,omitempty"`

	// **参数解释：** 起始ip。
	IpStart *string `json:"ip_start,omitempty"`

	// **参数解释：** 结束ip。
	IpEnd *string `json:"ip_end,omitempty"`

	// **参数解释：** 是否允许访问代码仓库。 - 0，表示禁止访问。 - 1，表示允许访问。
	ViewFlag *TrustedIpAddressDtoViewFlag `json:"view_flag,omitempty"`

	// **参数解释：** 是否允许下载代码。 - 0，表示禁止下载。 - 1，表示允许下载。
	DownloadFlag *TrustedIpAddressDtoDownloadFlag `json:"download_flag,omitempty"`

	// **参数解释：** 是否允许提交代码。 - 0，表示禁止提交。 - 1，表示允许提交。
	UploadFlag *TrustedIpAddressDtoUploadFlag `json:"upload_flag,omitempty"`

	// **参数解释：** 备注。 **取值范围：** 字符串长度不少于0，不超过200。
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 创建时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 排序。 - 0，表示默认规则。 - 1，表示自定义规则。
	OrderFlag *TrustedIpAddressDtoOrderFlag `json:"order_flag,omitempty"`
}

func (o TrustedIpAddressDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrustedIpAddressDto struct{}"
	}

	return strings.Join([]string{"TrustedIpAddressDto", string(data)}, " ")
}

type TrustedIpAddressDtoIpType struct {
	value int32
}

type TrustedIpAddressDtoIpTypeEnum struct {
	E_0 TrustedIpAddressDtoIpType
	E_1 TrustedIpAddressDtoIpType
	E_2 TrustedIpAddressDtoIpType
}

func GetTrustedIpAddressDtoIpTypeEnum() TrustedIpAddressDtoIpTypeEnum {
	return TrustedIpAddressDtoIpTypeEnum{
		E_0: TrustedIpAddressDtoIpType{
			value: 0,
		}, E_1: TrustedIpAddressDtoIpType{
			value: 1,
		}, E_2: TrustedIpAddressDtoIpType{
			value: 2,
		},
	}
}

func (c TrustedIpAddressDtoIpType) Value() int32 {
	return c.value
}

func (c TrustedIpAddressDtoIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrustedIpAddressDtoIpType) UnmarshalJSON(b []byte) error {
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

type TrustedIpAddressDtoViewFlag struct {
	value int32
}

type TrustedIpAddressDtoViewFlagEnum struct {
	E_0 TrustedIpAddressDtoViewFlag
	E_1 TrustedIpAddressDtoViewFlag
}

func GetTrustedIpAddressDtoViewFlagEnum() TrustedIpAddressDtoViewFlagEnum {
	return TrustedIpAddressDtoViewFlagEnum{
		E_0: TrustedIpAddressDtoViewFlag{
			value: 0,
		}, E_1: TrustedIpAddressDtoViewFlag{
			value: 1,
		},
	}
}

func (c TrustedIpAddressDtoViewFlag) Value() int32 {
	return c.value
}

func (c TrustedIpAddressDtoViewFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrustedIpAddressDtoViewFlag) UnmarshalJSON(b []byte) error {
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

type TrustedIpAddressDtoDownloadFlag struct {
	value int32
}

type TrustedIpAddressDtoDownloadFlagEnum struct {
	E_0 TrustedIpAddressDtoDownloadFlag
	E_1 TrustedIpAddressDtoDownloadFlag
}

func GetTrustedIpAddressDtoDownloadFlagEnum() TrustedIpAddressDtoDownloadFlagEnum {
	return TrustedIpAddressDtoDownloadFlagEnum{
		E_0: TrustedIpAddressDtoDownloadFlag{
			value: 0,
		}, E_1: TrustedIpAddressDtoDownloadFlag{
			value: 1,
		},
	}
}

func (c TrustedIpAddressDtoDownloadFlag) Value() int32 {
	return c.value
}

func (c TrustedIpAddressDtoDownloadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrustedIpAddressDtoDownloadFlag) UnmarshalJSON(b []byte) error {
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

type TrustedIpAddressDtoUploadFlag struct {
	value int32
}

type TrustedIpAddressDtoUploadFlagEnum struct {
	E_0 TrustedIpAddressDtoUploadFlag
	E_1 TrustedIpAddressDtoUploadFlag
}

func GetTrustedIpAddressDtoUploadFlagEnum() TrustedIpAddressDtoUploadFlagEnum {
	return TrustedIpAddressDtoUploadFlagEnum{
		E_0: TrustedIpAddressDtoUploadFlag{
			value: 0,
		}, E_1: TrustedIpAddressDtoUploadFlag{
			value: 1,
		},
	}
}

func (c TrustedIpAddressDtoUploadFlag) Value() int32 {
	return c.value
}

func (c TrustedIpAddressDtoUploadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrustedIpAddressDtoUploadFlag) UnmarshalJSON(b []byte) error {
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

type TrustedIpAddressDtoOrderFlag struct {
	value int32
}

type TrustedIpAddressDtoOrderFlagEnum struct {
	E_0 TrustedIpAddressDtoOrderFlag
	E_1 TrustedIpAddressDtoOrderFlag
}

func GetTrustedIpAddressDtoOrderFlagEnum() TrustedIpAddressDtoOrderFlagEnum {
	return TrustedIpAddressDtoOrderFlagEnum{
		E_0: TrustedIpAddressDtoOrderFlag{
			value: 0,
		}, E_1: TrustedIpAddressDtoOrderFlag{
			value: 1,
		},
	}
}

func (c TrustedIpAddressDtoOrderFlag) Value() int32 {
	return c.value
}

func (c TrustedIpAddressDtoOrderFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TrustedIpAddressDtoOrderFlag) UnmarshalJSON(b []byte) error {
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
