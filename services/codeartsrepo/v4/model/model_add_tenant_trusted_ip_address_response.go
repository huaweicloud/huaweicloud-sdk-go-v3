package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddTenantTrustedIpAddressResponse Response Object
type AddTenantTrustedIpAddressResponse struct {

	// **参数解释：** ip白名单id。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 用户id。
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释：** 租户id。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** ip范围。
	IpRange *string `json:"ip_range,omitempty"`

	// **参数解释：** 格式类型。 - 0，表示指定IP。 - 1，表示ip范围。 - 2，表示CIDR。
	IpType *AddTenantTrustedIpAddressResponseIpType `json:"ip_type,omitempty"`

	// **参数解释：** 起始ip。
	IpStart *string `json:"ip_start,omitempty"`

	// **参数解释：** 结束ip。
	IpEnd *string `json:"ip_end,omitempty"`

	// **参数解释：** 是否允许访问代码仓库。 - 0，表示禁止访问。 - 1，表示允许访问。
	ViewFlag *AddTenantTrustedIpAddressResponseViewFlag `json:"view_flag,omitempty"`

	// **参数解释：** 是否允许下载代码。 - 0，表示禁止下载。 - 1，表示允许下载。
	DownloadFlag *AddTenantTrustedIpAddressResponseDownloadFlag `json:"download_flag,omitempty"`

	// **参数解释：** 是否允许提交代码。 - 0，表示禁止提交。 - 1，表示允许提交。
	UploadFlag *AddTenantTrustedIpAddressResponseUploadFlag `json:"upload_flag,omitempty"`

	// **参数解释：** 备注。 **取值范围：** 字符串长度不少于0，不超过200。
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 创建时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 排序。 - 0，表示默认规则。 - 1，表示自定义规则。
	OrderFlag      *AddTenantTrustedIpAddressResponseOrderFlag `json:"order_flag,omitempty"`
	HttpStatusCode int                                         `json:"-"`
}

func (o AddTenantTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTenantTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"AddTenantTrustedIpAddressResponse", string(data)}, " ")
}

type AddTenantTrustedIpAddressResponseIpType struct {
	value int32
}

type AddTenantTrustedIpAddressResponseIpTypeEnum struct {
	E_0 AddTenantTrustedIpAddressResponseIpType
	E_1 AddTenantTrustedIpAddressResponseIpType
	E_2 AddTenantTrustedIpAddressResponseIpType
}

func GetAddTenantTrustedIpAddressResponseIpTypeEnum() AddTenantTrustedIpAddressResponseIpTypeEnum {
	return AddTenantTrustedIpAddressResponseIpTypeEnum{
		E_0: AddTenantTrustedIpAddressResponseIpType{
			value: 0,
		}, E_1: AddTenantTrustedIpAddressResponseIpType{
			value: 1,
		}, E_2: AddTenantTrustedIpAddressResponseIpType{
			value: 2,
		},
	}
}

func (c AddTenantTrustedIpAddressResponseIpType) Value() int32 {
	return c.value
}

func (c AddTenantTrustedIpAddressResponseIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTenantTrustedIpAddressResponseIpType) UnmarshalJSON(b []byte) error {
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

type AddTenantTrustedIpAddressResponseViewFlag struct {
	value int32
}

type AddTenantTrustedIpAddressResponseViewFlagEnum struct {
	E_0 AddTenantTrustedIpAddressResponseViewFlag
	E_1 AddTenantTrustedIpAddressResponseViewFlag
}

func GetAddTenantTrustedIpAddressResponseViewFlagEnum() AddTenantTrustedIpAddressResponseViewFlagEnum {
	return AddTenantTrustedIpAddressResponseViewFlagEnum{
		E_0: AddTenantTrustedIpAddressResponseViewFlag{
			value: 0,
		}, E_1: AddTenantTrustedIpAddressResponseViewFlag{
			value: 1,
		},
	}
}

func (c AddTenantTrustedIpAddressResponseViewFlag) Value() int32 {
	return c.value
}

func (c AddTenantTrustedIpAddressResponseViewFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTenantTrustedIpAddressResponseViewFlag) UnmarshalJSON(b []byte) error {
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

type AddTenantTrustedIpAddressResponseDownloadFlag struct {
	value int32
}

type AddTenantTrustedIpAddressResponseDownloadFlagEnum struct {
	E_0 AddTenantTrustedIpAddressResponseDownloadFlag
	E_1 AddTenantTrustedIpAddressResponseDownloadFlag
}

func GetAddTenantTrustedIpAddressResponseDownloadFlagEnum() AddTenantTrustedIpAddressResponseDownloadFlagEnum {
	return AddTenantTrustedIpAddressResponseDownloadFlagEnum{
		E_0: AddTenantTrustedIpAddressResponseDownloadFlag{
			value: 0,
		}, E_1: AddTenantTrustedIpAddressResponseDownloadFlag{
			value: 1,
		},
	}
}

func (c AddTenantTrustedIpAddressResponseDownloadFlag) Value() int32 {
	return c.value
}

func (c AddTenantTrustedIpAddressResponseDownloadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTenantTrustedIpAddressResponseDownloadFlag) UnmarshalJSON(b []byte) error {
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

type AddTenantTrustedIpAddressResponseUploadFlag struct {
	value int32
}

type AddTenantTrustedIpAddressResponseUploadFlagEnum struct {
	E_0 AddTenantTrustedIpAddressResponseUploadFlag
	E_1 AddTenantTrustedIpAddressResponseUploadFlag
}

func GetAddTenantTrustedIpAddressResponseUploadFlagEnum() AddTenantTrustedIpAddressResponseUploadFlagEnum {
	return AddTenantTrustedIpAddressResponseUploadFlagEnum{
		E_0: AddTenantTrustedIpAddressResponseUploadFlag{
			value: 0,
		}, E_1: AddTenantTrustedIpAddressResponseUploadFlag{
			value: 1,
		},
	}
}

func (c AddTenantTrustedIpAddressResponseUploadFlag) Value() int32 {
	return c.value
}

func (c AddTenantTrustedIpAddressResponseUploadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTenantTrustedIpAddressResponseUploadFlag) UnmarshalJSON(b []byte) error {
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

type AddTenantTrustedIpAddressResponseOrderFlag struct {
	value int32
}

type AddTenantTrustedIpAddressResponseOrderFlagEnum struct {
	E_0 AddTenantTrustedIpAddressResponseOrderFlag
	E_1 AddTenantTrustedIpAddressResponseOrderFlag
}

func GetAddTenantTrustedIpAddressResponseOrderFlagEnum() AddTenantTrustedIpAddressResponseOrderFlagEnum {
	return AddTenantTrustedIpAddressResponseOrderFlagEnum{
		E_0: AddTenantTrustedIpAddressResponseOrderFlag{
			value: 0,
		}, E_1: AddTenantTrustedIpAddressResponseOrderFlag{
			value: 1,
		},
	}
}

func (c AddTenantTrustedIpAddressResponseOrderFlag) Value() int32 {
	return c.value
}

func (c AddTenantTrustedIpAddressResponseOrderFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTenantTrustedIpAddressResponseOrderFlag) UnmarshalJSON(b []byte) error {
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
