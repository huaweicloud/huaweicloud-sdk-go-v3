package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SignatureRequest 短信签名请求体。
type SignatureRequest struct {

	// 签名名称。
	SignatureName string `json:"signature_name"`

	// 签名类型。PROMOTION_TYPE：营销类，NOTIFY_TYPE：通知类。
	SignatureType SignatureRequestSignatureType `json:"signature_type"`

	// 短信应用ID。
	AppId string `json:"app_id"`

	// 申请说明。
	ApplyDesc *string `json:"apply_desc,omitempty"`

	// 营业执照文件ID。调用上传申请文件接口获取。
	FileId string `json:"file_id"`

	// 签名来源。0：企事业单位的全称或简称，1：工信部备案网站的全称或简称，2：APP应用的全称或简称，3：公众号或小程序的全称或简称，4：电商平台店铺名的全称或简称，5：商标名的全称或简称。
	SignatureSource int32 `json:"signature_source"`

	// 是否涉及第三方权益。若为yes，则还需要传入授权委托书。yes：涉及，no：不涉及。
	IsInvolvedThird SignatureRequestIsInvolvedThird `json:"is_involved_third"`

	// 授权委托书文件ID。调用上传申请文件接口获取。
	PowerAttorneyFileId *string `json:"power_attorney_file_id,omitempty"`

	// 签名来源标题内容，填写内容需和签名来源对应，具体如下。 - 网站域名：仅当“签名来源”为“工信部备案网站的全称或简称”时，需要输入工信部备案网站域名，如huawei.com。 - APP应用下载地址：仅当“签名来源”为“APP应用的全称或简称”时，需要填写。 - 公众号或者小程序的全称：仅当“签名来源”为“公众号或小程序的全称或简称”时，需要填写。 - 电商平台店铺地址：仅当“签名来源”为“电商平台店铺名的全称或简称”时，需要填写。
	SourceTitleContent *string `json:"source_title_content,omitempty"`
}

func (o SignatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SignatureRequest struct{}"
	}

	return strings.Join([]string{"SignatureRequest", string(data)}, " ")
}

type SignatureRequestSignatureType struct {
	value string
}

type SignatureRequestSignatureTypeEnum struct {
	PROMOTION_TYPE SignatureRequestSignatureType
	NOTIFY_TYPE    SignatureRequestSignatureType
}

func GetSignatureRequestSignatureTypeEnum() SignatureRequestSignatureTypeEnum {
	return SignatureRequestSignatureTypeEnum{
		PROMOTION_TYPE: SignatureRequestSignatureType{
			value: "PROMOTION_TYPE",
		},
		NOTIFY_TYPE: SignatureRequestSignatureType{
			value: "NOTIFY_TYPE",
		},
	}
}

func (c SignatureRequestSignatureType) Value() string {
	return c.value
}

func (c SignatureRequestSignatureType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignatureRequestSignatureType) UnmarshalJSON(b []byte) error {
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

type SignatureRequestIsInvolvedThird struct {
	value string
}

type SignatureRequestIsInvolvedThirdEnum struct {
	YES SignatureRequestIsInvolvedThird
	NO  SignatureRequestIsInvolvedThird
}

func GetSignatureRequestIsInvolvedThirdEnum() SignatureRequestIsInvolvedThirdEnum {
	return SignatureRequestIsInvolvedThirdEnum{
		YES: SignatureRequestIsInvolvedThird{
			value: "yes",
		},
		NO: SignatureRequestIsInvolvedThird{
			value: "no",
		},
	}
}

func (c SignatureRequestIsInvolvedThird) Value() string {
	return c.value
}

func (c SignatureRequestIsInvolvedThird) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SignatureRequestIsInvolvedThird) UnmarshalJSON(b []byte) error {
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
