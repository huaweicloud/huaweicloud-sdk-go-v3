package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AccessoryJsonFormatReq struct {

	// 文件名称
	AccessoryName string `json:"accessory_name"`

	// 文件来源，创建工单附件传incident，留言附件传message
	AccessoryFrom AccessoryJsonFormatReqAccessoryFrom `json:"accessory_from"`

	// 上传类型，默认为0，markdown模式为1
	UploadType *int32 `json:"upload_type,omitempty"`

	// 文件内容，Base64格式
	AccessoryData string `json:"accessory_data"`
}

func (o AccessoryJsonFormatReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessoryJsonFormatReq struct{}"
	}

	return strings.Join([]string{"AccessoryJsonFormatReq", string(data)}, " ")
}

type AccessoryJsonFormatReqAccessoryFrom struct {
	value string
}

type AccessoryJsonFormatReqAccessoryFromEnum struct {
	INCIDENT AccessoryJsonFormatReqAccessoryFrom
	MESSAGE  AccessoryJsonFormatReqAccessoryFrom
}

func GetAccessoryJsonFormatReqAccessoryFromEnum() AccessoryJsonFormatReqAccessoryFromEnum {
	return AccessoryJsonFormatReqAccessoryFromEnum{
		INCIDENT: AccessoryJsonFormatReqAccessoryFrom{
			value: "incident",
		},
		MESSAGE: AccessoryJsonFormatReqAccessoryFrom{
			value: "message",
		},
	}
}

func (c AccessoryJsonFormatReqAccessoryFrom) Value() string {
	return c.value
}

func (c AccessoryJsonFormatReqAccessoryFrom) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessoryJsonFormatReqAccessoryFrom) UnmarshalJSON(b []byte) error {
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
