package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTenantTrustedIpAddressResponse Response Object
type UpdateTenantTrustedIpAddressResponse struct {

	// **参数解释：** ip白名单id。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 用户id。
	UserId *int32 `json:"user_id,omitempty"`

	// **参数解释：** 租户id。
	DomainId *string `json:"domain_id,omitempty"`

	// **参数解释：** ip范围。
	IpRange *string `json:"ip_range,omitempty"`

	// **参数解释：** 格式类型。 - 0，表示指定IP。 - 1，表示ip范围。 - 2，表示CIDR。
	IpType *UpdateTenantTrustedIpAddressResponseIpType `json:"ip_type,omitempty"`

	// **参数解释：** 起始ip。
	IpStart *string `json:"ip_start,omitempty"`

	// **参数解释：** 结束ip。
	IpEnd *string `json:"ip_end,omitempty"`

	// **参数解释：** 是否允许访问代码仓库。 - 0，表示禁止访问。 - 1，表示允许访问。
	ViewFlag *UpdateTenantTrustedIpAddressResponseViewFlag `json:"view_flag,omitempty"`

	// **参数解释：** 是否允许下载代码。 - 0，表示禁止下载。 - 1，表示允许下载。
	DownloadFlag *UpdateTenantTrustedIpAddressResponseDownloadFlag `json:"download_flag,omitempty"`

	// **参数解释：** 是否允许提交代码。 - 0，表示禁止提交。 - 1，表示允许提交。
	UploadFlag *UpdateTenantTrustedIpAddressResponseUploadFlag `json:"upload_flag,omitempty"`

	// **参数解释：** 备注。 **取值范围：** 字符串长度不少于0，不超过200。
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 创建时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 排序。 - 0，表示默认规则。 - 1，表示自定义规则。
	OrderFlag      *UpdateTenantTrustedIpAddressResponseOrderFlag `json:"order_flag,omitempty"`
	HttpStatusCode int                                            `json:"-"`
}

func (o UpdateTenantTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTenantTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"UpdateTenantTrustedIpAddressResponse", string(data)}, " ")
}

type UpdateTenantTrustedIpAddressResponseIpType struct {
	value int32
}

type UpdateTenantTrustedIpAddressResponseIpTypeEnum struct {
	E_0 UpdateTenantTrustedIpAddressResponseIpType
	E_1 UpdateTenantTrustedIpAddressResponseIpType
	E_2 UpdateTenantTrustedIpAddressResponseIpType
}

func GetUpdateTenantTrustedIpAddressResponseIpTypeEnum() UpdateTenantTrustedIpAddressResponseIpTypeEnum {
	return UpdateTenantTrustedIpAddressResponseIpTypeEnum{
		E_0: UpdateTenantTrustedIpAddressResponseIpType{
			value: 0,
		}, E_1: UpdateTenantTrustedIpAddressResponseIpType{
			value: 1,
		}, E_2: UpdateTenantTrustedIpAddressResponseIpType{
			value: 2,
		},
	}
}

func (c UpdateTenantTrustedIpAddressResponseIpType) Value() int32 {
	return c.value
}

func (c UpdateTenantTrustedIpAddressResponseIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTenantTrustedIpAddressResponseIpType) UnmarshalJSON(b []byte) error {
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

type UpdateTenantTrustedIpAddressResponseViewFlag struct {
	value int32
}

type UpdateTenantTrustedIpAddressResponseViewFlagEnum struct {
	E_0 UpdateTenantTrustedIpAddressResponseViewFlag
	E_1 UpdateTenantTrustedIpAddressResponseViewFlag
}

func GetUpdateTenantTrustedIpAddressResponseViewFlagEnum() UpdateTenantTrustedIpAddressResponseViewFlagEnum {
	return UpdateTenantTrustedIpAddressResponseViewFlagEnum{
		E_0: UpdateTenantTrustedIpAddressResponseViewFlag{
			value: 0,
		}, E_1: UpdateTenantTrustedIpAddressResponseViewFlag{
			value: 1,
		},
	}
}

func (c UpdateTenantTrustedIpAddressResponseViewFlag) Value() int32 {
	return c.value
}

func (c UpdateTenantTrustedIpAddressResponseViewFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTenantTrustedIpAddressResponseViewFlag) UnmarshalJSON(b []byte) error {
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

type UpdateTenantTrustedIpAddressResponseDownloadFlag struct {
	value int32
}

type UpdateTenantTrustedIpAddressResponseDownloadFlagEnum struct {
	E_0 UpdateTenantTrustedIpAddressResponseDownloadFlag
	E_1 UpdateTenantTrustedIpAddressResponseDownloadFlag
}

func GetUpdateTenantTrustedIpAddressResponseDownloadFlagEnum() UpdateTenantTrustedIpAddressResponseDownloadFlagEnum {
	return UpdateTenantTrustedIpAddressResponseDownloadFlagEnum{
		E_0: UpdateTenantTrustedIpAddressResponseDownloadFlag{
			value: 0,
		}, E_1: UpdateTenantTrustedIpAddressResponseDownloadFlag{
			value: 1,
		},
	}
}

func (c UpdateTenantTrustedIpAddressResponseDownloadFlag) Value() int32 {
	return c.value
}

func (c UpdateTenantTrustedIpAddressResponseDownloadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTenantTrustedIpAddressResponseDownloadFlag) UnmarshalJSON(b []byte) error {
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

type UpdateTenantTrustedIpAddressResponseUploadFlag struct {
	value int32
}

type UpdateTenantTrustedIpAddressResponseUploadFlagEnum struct {
	E_0 UpdateTenantTrustedIpAddressResponseUploadFlag
	E_1 UpdateTenantTrustedIpAddressResponseUploadFlag
}

func GetUpdateTenantTrustedIpAddressResponseUploadFlagEnum() UpdateTenantTrustedIpAddressResponseUploadFlagEnum {
	return UpdateTenantTrustedIpAddressResponseUploadFlagEnum{
		E_0: UpdateTenantTrustedIpAddressResponseUploadFlag{
			value: 0,
		}, E_1: UpdateTenantTrustedIpAddressResponseUploadFlag{
			value: 1,
		},
	}
}

func (c UpdateTenantTrustedIpAddressResponseUploadFlag) Value() int32 {
	return c.value
}

func (c UpdateTenantTrustedIpAddressResponseUploadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTenantTrustedIpAddressResponseUploadFlag) UnmarshalJSON(b []byte) error {
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

type UpdateTenantTrustedIpAddressResponseOrderFlag struct {
	value int32
}

type UpdateTenantTrustedIpAddressResponseOrderFlagEnum struct {
	E_0 UpdateTenantTrustedIpAddressResponseOrderFlag
	E_1 UpdateTenantTrustedIpAddressResponseOrderFlag
}

func GetUpdateTenantTrustedIpAddressResponseOrderFlagEnum() UpdateTenantTrustedIpAddressResponseOrderFlagEnum {
	return UpdateTenantTrustedIpAddressResponseOrderFlagEnum{
		E_0: UpdateTenantTrustedIpAddressResponseOrderFlag{
			value: 0,
		}, E_1: UpdateTenantTrustedIpAddressResponseOrderFlag{
			value: 1,
		},
	}
}

func (c UpdateTenantTrustedIpAddressResponseOrderFlag) Value() int32 {
	return c.value
}

func (c UpdateTenantTrustedIpAddressResponseOrderFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTenantTrustedIpAddressResponseOrderFlag) UnmarshalJSON(b []byte) error {
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
