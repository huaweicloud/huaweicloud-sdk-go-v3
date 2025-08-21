package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AddTrustedIpAddressRequestBody **参数解释：** 添加ip白名单请求体。
type AddTrustedIpAddressRequestBody struct {

	// **参数解释：** 格式类型。 - 0，表示指定ip。 - 1，表示ip范围。 - 2，表示CIDR。
	IpType *AddTrustedIpAddressRequestBodyIpType `json:"ip_type,omitempty"`

	// **参数解释：** 起始ip。
	IpStart *string `json:"ip_start,omitempty"`

	// **参数解释：** 结束ip。
	IpEnd *string `json:"ip_end,omitempty"`

	// **参数解释：** 是否允许访问代码仓库。 - 0，表示禁止访问。 - 1，表示允许访问。
	ViewFlag *AddTrustedIpAddressRequestBodyViewFlag `json:"view_flag,omitempty"`

	// **参数解释：** 是否允许下载代码。 - 0，表示禁止下载。 - 1，表示允许下载。
	DownloadFlag *AddTrustedIpAddressRequestBodyDownloadFlag `json:"download_flag,omitempty"`

	// **参数解释：** 是否允许提交代码。 - 0，表示禁止提交。 - 1，表示允许提交。
	UploadFlag *AddTrustedIpAddressRequestBodyUploadFlag `json:"upload_flag,omitempty"`

	// **参数解释：** 备注。 **取值范围：** 字符串长度不少于0，不超过200。
	Remark *string `json:"remark,omitempty"`
}

func (o AddTrustedIpAddressRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddTrustedIpAddressRequestBody struct{}"
	}

	return strings.Join([]string{"AddTrustedIpAddressRequestBody", string(data)}, " ")
}

type AddTrustedIpAddressRequestBodyIpType struct {
	value int32
}

type AddTrustedIpAddressRequestBodyIpTypeEnum struct {
	E_0 AddTrustedIpAddressRequestBodyIpType
	E_1 AddTrustedIpAddressRequestBodyIpType
	E_2 AddTrustedIpAddressRequestBodyIpType
}

func GetAddTrustedIpAddressRequestBodyIpTypeEnum() AddTrustedIpAddressRequestBodyIpTypeEnum {
	return AddTrustedIpAddressRequestBodyIpTypeEnum{
		E_0: AddTrustedIpAddressRequestBodyIpType{
			value: 0,
		}, E_1: AddTrustedIpAddressRequestBodyIpType{
			value: 1,
		}, E_2: AddTrustedIpAddressRequestBodyIpType{
			value: 2,
		},
	}
}

func (c AddTrustedIpAddressRequestBodyIpType) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressRequestBodyIpType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressRequestBodyIpType) UnmarshalJSON(b []byte) error {
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

type AddTrustedIpAddressRequestBodyViewFlag struct {
	value int32
}

type AddTrustedIpAddressRequestBodyViewFlagEnum struct {
	E_0 AddTrustedIpAddressRequestBodyViewFlag
	E_1 AddTrustedIpAddressRequestBodyViewFlag
}

func GetAddTrustedIpAddressRequestBodyViewFlagEnum() AddTrustedIpAddressRequestBodyViewFlagEnum {
	return AddTrustedIpAddressRequestBodyViewFlagEnum{
		E_0: AddTrustedIpAddressRequestBodyViewFlag{
			value: 0,
		}, E_1: AddTrustedIpAddressRequestBodyViewFlag{
			value: 1,
		},
	}
}

func (c AddTrustedIpAddressRequestBodyViewFlag) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressRequestBodyViewFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressRequestBodyViewFlag) UnmarshalJSON(b []byte) error {
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

type AddTrustedIpAddressRequestBodyDownloadFlag struct {
	value int32
}

type AddTrustedIpAddressRequestBodyDownloadFlagEnum struct {
	E_0 AddTrustedIpAddressRequestBodyDownloadFlag
	E_1 AddTrustedIpAddressRequestBodyDownloadFlag
}

func GetAddTrustedIpAddressRequestBodyDownloadFlagEnum() AddTrustedIpAddressRequestBodyDownloadFlagEnum {
	return AddTrustedIpAddressRequestBodyDownloadFlagEnum{
		E_0: AddTrustedIpAddressRequestBodyDownloadFlag{
			value: 0,
		}, E_1: AddTrustedIpAddressRequestBodyDownloadFlag{
			value: 1,
		},
	}
}

func (c AddTrustedIpAddressRequestBodyDownloadFlag) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressRequestBodyDownloadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressRequestBodyDownloadFlag) UnmarshalJSON(b []byte) error {
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

type AddTrustedIpAddressRequestBodyUploadFlag struct {
	value int32
}

type AddTrustedIpAddressRequestBodyUploadFlagEnum struct {
	E_0 AddTrustedIpAddressRequestBodyUploadFlag
	E_1 AddTrustedIpAddressRequestBodyUploadFlag
}

func GetAddTrustedIpAddressRequestBodyUploadFlagEnum() AddTrustedIpAddressRequestBodyUploadFlagEnum {
	return AddTrustedIpAddressRequestBodyUploadFlagEnum{
		E_0: AddTrustedIpAddressRequestBodyUploadFlag{
			value: 0,
		}, E_1: AddTrustedIpAddressRequestBodyUploadFlag{
			value: 1,
		},
	}
}

func (c AddTrustedIpAddressRequestBodyUploadFlag) Value() int32 {
	return c.value
}

func (c AddTrustedIpAddressRequestBodyUploadFlag) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AddTrustedIpAddressRequestBodyUploadFlag) UnmarshalJSON(b []byte) error {
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
