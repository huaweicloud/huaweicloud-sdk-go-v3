package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTrustedIpAddressResponse Response Object
type UpdateTrustedIpAddressResponse struct {

	// **参数解释：** 白名单id。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** ip范围。
	IpRange *string `json:"ip_range,omitempty"`

	// **参数解释：** 格式类型。 - 0，表示指定ip。 - 1，表示ip范围。 - 2，表示CIDR。
	IpType *UpdateTrustedIpAddressResponseIpType `json:"ip_type,omitempty"`

	// **参数解释：** 起始ip。
	IpStart *string `json:"ip_start,omitempty"`

	// **参数解释：** 结束ip。
	IpEnd *string `json:"ip_end,omitempty"`

	// **参数解释：** 是否允许访问代码仓库。 - 0，表示禁止访问。 - 1，表示允许访问。
	ViewFlag *UpdateTrustedIpAddressResponseViewFlag `json:"view_flag,omitempty"`

	// **参数解释：** 是否允许下载代码。 - 0，表示禁止下载。 - 1，表示允许下载。
	DownloadFlag *UpdateTrustedIpAddressResponseDownloadFlag `json:"download_flag,omitempty"`

	// **参数解释：** 是否允许提交代码。 - 0，表示禁止提交。 - 1，表示允许提交。
	UploadFlag *UpdateTrustedIpAddressResponseUploadFlag `json:"upload_flag,omitempty"`

	// **参数解释：** 备注。 **取值范围：** 字符串长度不少于0，不超过200。
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 创建时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 排序。 - 0，表示默认规则。 - 1，表示自定义规则。
	OrderFlag      *UpdateTrustedIpAddressResponseOrderFlag `json:"order_flag,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o UpdateTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"UpdateTrustedIpAddressResponse", string(data)}, " ")
}

type UpdateTrustedIpAddressResponseIpType struct {
	value int32
}

type UpdateTrustedIpAddressResponseIpTypeEnum struct {
	E_0 UpdateTrustedIpAddressResponseIpType
	E_1 UpdateTrustedIpAddressResponseIpType
	E_2 UpdateTrustedIpAddressResponseIpType
}

func GetUpdateTrustedIpAddressResponseIpTypeEnum() UpdateTrustedIpAddressResponseIpTypeEnum {
	return UpdateTrustedIpAddressResponseIpTypeEnum{
		E_0: UpdateTrustedIpAddressResponseIpType{
			value: 0,
		}, E_1: UpdateTrustedIpAddressResponseIpType{
			value: 1,
		}, E_2: UpdateTrustedIpAddressResponseIpType{
			value: 2,
		},
	}
}

func (c UpdateTrustedIpAddressResponseIpType) Value() int32 {
	return c.value
}

func (c UpdateTrustedIpAddressResponseIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrustedIpAddressResponseIpType) UnmarshalJSON(b []byte) error {
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

type UpdateTrustedIpAddressResponseViewFlag struct {
	value int32
}

type UpdateTrustedIpAddressResponseViewFlagEnum struct {
	E_0 UpdateTrustedIpAddressResponseViewFlag
	E_1 UpdateTrustedIpAddressResponseViewFlag
}

func GetUpdateTrustedIpAddressResponseViewFlagEnum() UpdateTrustedIpAddressResponseViewFlagEnum {
	return UpdateTrustedIpAddressResponseViewFlagEnum{
		E_0: UpdateTrustedIpAddressResponseViewFlag{
			value: 0,
		}, E_1: UpdateTrustedIpAddressResponseViewFlag{
			value: 1,
		},
	}
}

func (c UpdateTrustedIpAddressResponseViewFlag) Value() int32 {
	return c.value
}

func (c UpdateTrustedIpAddressResponseViewFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrustedIpAddressResponseViewFlag) UnmarshalJSON(b []byte) error {
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

type UpdateTrustedIpAddressResponseDownloadFlag struct {
	value int32
}

type UpdateTrustedIpAddressResponseDownloadFlagEnum struct {
	E_0 UpdateTrustedIpAddressResponseDownloadFlag
	E_1 UpdateTrustedIpAddressResponseDownloadFlag
}

func GetUpdateTrustedIpAddressResponseDownloadFlagEnum() UpdateTrustedIpAddressResponseDownloadFlagEnum {
	return UpdateTrustedIpAddressResponseDownloadFlagEnum{
		E_0: UpdateTrustedIpAddressResponseDownloadFlag{
			value: 0,
		}, E_1: UpdateTrustedIpAddressResponseDownloadFlag{
			value: 1,
		},
	}
}

func (c UpdateTrustedIpAddressResponseDownloadFlag) Value() int32 {
	return c.value
}

func (c UpdateTrustedIpAddressResponseDownloadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrustedIpAddressResponseDownloadFlag) UnmarshalJSON(b []byte) error {
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

type UpdateTrustedIpAddressResponseUploadFlag struct {
	value int32
}

type UpdateTrustedIpAddressResponseUploadFlagEnum struct {
	E_0 UpdateTrustedIpAddressResponseUploadFlag
	E_1 UpdateTrustedIpAddressResponseUploadFlag
}

func GetUpdateTrustedIpAddressResponseUploadFlagEnum() UpdateTrustedIpAddressResponseUploadFlagEnum {
	return UpdateTrustedIpAddressResponseUploadFlagEnum{
		E_0: UpdateTrustedIpAddressResponseUploadFlag{
			value: 0,
		}, E_1: UpdateTrustedIpAddressResponseUploadFlag{
			value: 1,
		},
	}
}

func (c UpdateTrustedIpAddressResponseUploadFlag) Value() int32 {
	return c.value
}

func (c UpdateTrustedIpAddressResponseUploadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrustedIpAddressResponseUploadFlag) UnmarshalJSON(b []byte) error {
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

type UpdateTrustedIpAddressResponseOrderFlag struct {
	value int32
}

type UpdateTrustedIpAddressResponseOrderFlagEnum struct {
	E_0 UpdateTrustedIpAddressResponseOrderFlag
	E_1 UpdateTrustedIpAddressResponseOrderFlag
}

func GetUpdateTrustedIpAddressResponseOrderFlagEnum() UpdateTrustedIpAddressResponseOrderFlagEnum {
	return UpdateTrustedIpAddressResponseOrderFlagEnum{
		E_0: UpdateTrustedIpAddressResponseOrderFlag{
			value: 0,
		}, E_1: UpdateTrustedIpAddressResponseOrderFlag{
			value: 1,
		},
	}
}

func (c UpdateTrustedIpAddressResponseOrderFlag) Value() int32 {
	return c.value
}

func (c UpdateTrustedIpAddressResponseOrderFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTrustedIpAddressResponseOrderFlag) UnmarshalJSON(b []byte) error {
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
