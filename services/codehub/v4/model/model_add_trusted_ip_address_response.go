package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddTrustedIpAddressResponse Response Object
type AddTrustedIpAddressResponse struct {

	// **参数解释：** 白名单id。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 仓库id。
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// **参数解释：** ip范围。
	IpRange *string `json:"ip_range,omitempty"`

	// **参数解释：** 格式类型。 - 0，表示指定ip。 - 1，表示ip范围。 - 2，表示CIDR。
	IpType *AddTrustedIpAddressResponseIpType `json:"ip_type,omitempty"`

	// **参数解释：** 起始ip。
	IpStart *string `json:"ip_start,omitempty"`

	// **参数解释：** 结束ip。
	IpEnd *string `json:"ip_end,omitempty"`

	// **参数解释：** 是否允许访问代码仓库。 - 0，表示禁止访问。 - 1，表示允许访问。
	ViewFlag *AddTrustedIpAddressResponseViewFlag `json:"view_flag,omitempty"`

	// **参数解释：** 是否允许下载代码。 - 0，表示禁止下载。 - 1，表示允许下载。
	DownloadFlag *AddTrustedIpAddressResponseDownloadFlag `json:"download_flag,omitempty"`

	// **参数解释：** 是否允许提交代码。 - 0，表示禁止提交。 - 1，表示允许提交。
	UploadFlag *AddTrustedIpAddressResponseUploadFlag `json:"upload_flag,omitempty"`

	// **参数解释：** 备注。 **取值范围：** 字符串长度不少于0，不超过200。
	Remark *string `json:"remark,omitempty"`

	// **参数解释：** 创建时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 更新时间。 **参数解释：** MMM dd, yyyy hh:mm:ss a
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 排序。 - 0，表示默认规则。 - 1，表示自定义规则。
	OrderFlag      *AddTrustedIpAddressResponseOrderFlag `json:"order_flag,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o AddTrustedIpAddressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTrustedIpAddressResponse struct{}"
	}

	return strings.Join([]string{"AddTrustedIpAddressResponse", string(data)}, " ")
}

type AddTrustedIpAddressResponseIpType struct {
	value int32
}

type AddTrustedIpAddressResponseIpTypeEnum struct {
	E_0 AddTrustedIpAddressResponseIpType
	E_1 AddTrustedIpAddressResponseIpType
	E_2 AddTrustedIpAddressResponseIpType
}

func GetAddTrustedIpAddressResponseIpTypeEnum() AddTrustedIpAddressResponseIpTypeEnum {
	return AddTrustedIpAddressResponseIpTypeEnum{
		E_0: AddTrustedIpAddressResponseIpType{
			value: 0,
		}, E_1: AddTrustedIpAddressResponseIpType{
			value: 1,
		}, E_2: AddTrustedIpAddressResponseIpType{
			value: 2,
		},
	}
}

func (c AddTrustedIpAddressResponseIpType) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressResponseIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressResponseIpType) UnmarshalJSON(b []byte) error {
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

type AddTrustedIpAddressResponseViewFlag struct {
	value int32
}

type AddTrustedIpAddressResponseViewFlagEnum struct {
	E_0 AddTrustedIpAddressResponseViewFlag
	E_1 AddTrustedIpAddressResponseViewFlag
}

func GetAddTrustedIpAddressResponseViewFlagEnum() AddTrustedIpAddressResponseViewFlagEnum {
	return AddTrustedIpAddressResponseViewFlagEnum{
		E_0: AddTrustedIpAddressResponseViewFlag{
			value: 0,
		}, E_1: AddTrustedIpAddressResponseViewFlag{
			value: 1,
		},
	}
}

func (c AddTrustedIpAddressResponseViewFlag) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressResponseViewFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressResponseViewFlag) UnmarshalJSON(b []byte) error {
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

type AddTrustedIpAddressResponseDownloadFlag struct {
	value int32
}

type AddTrustedIpAddressResponseDownloadFlagEnum struct {
	E_0 AddTrustedIpAddressResponseDownloadFlag
	E_1 AddTrustedIpAddressResponseDownloadFlag
}

func GetAddTrustedIpAddressResponseDownloadFlagEnum() AddTrustedIpAddressResponseDownloadFlagEnum {
	return AddTrustedIpAddressResponseDownloadFlagEnum{
		E_0: AddTrustedIpAddressResponseDownloadFlag{
			value: 0,
		}, E_1: AddTrustedIpAddressResponseDownloadFlag{
			value: 1,
		},
	}
}

func (c AddTrustedIpAddressResponseDownloadFlag) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressResponseDownloadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressResponseDownloadFlag) UnmarshalJSON(b []byte) error {
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

type AddTrustedIpAddressResponseUploadFlag struct {
	value int32
}

type AddTrustedIpAddressResponseUploadFlagEnum struct {
	E_0 AddTrustedIpAddressResponseUploadFlag
	E_1 AddTrustedIpAddressResponseUploadFlag
}

func GetAddTrustedIpAddressResponseUploadFlagEnum() AddTrustedIpAddressResponseUploadFlagEnum {
	return AddTrustedIpAddressResponseUploadFlagEnum{
		E_0: AddTrustedIpAddressResponseUploadFlag{
			value: 0,
		}, E_1: AddTrustedIpAddressResponseUploadFlag{
			value: 1,
		},
	}
}

func (c AddTrustedIpAddressResponseUploadFlag) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressResponseUploadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressResponseUploadFlag) UnmarshalJSON(b []byte) error {
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

type AddTrustedIpAddressResponseOrderFlag struct {
	value int32
}

type AddTrustedIpAddressResponseOrderFlagEnum struct {
	E_0 AddTrustedIpAddressResponseOrderFlag
	E_1 AddTrustedIpAddressResponseOrderFlag
}

func GetAddTrustedIpAddressResponseOrderFlagEnum() AddTrustedIpAddressResponseOrderFlagEnum {
	return AddTrustedIpAddressResponseOrderFlagEnum{
		E_0: AddTrustedIpAddressResponseOrderFlag{
			value: 0,
		}, E_1: AddTrustedIpAddressResponseOrderFlag{
			value: 1,
		},
	}
}

func (c AddTrustedIpAddressResponseOrderFlag) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressResponseOrderFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressResponseOrderFlag) UnmarshalJSON(b []byte) error {
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
