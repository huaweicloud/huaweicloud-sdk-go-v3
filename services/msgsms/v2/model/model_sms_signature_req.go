package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsSignatureReq struct {

	// 应用主键ID
	AppId string `json:"app_id"`

	// 申请说明
	ApplyDesc *string `json:"apply_desc,omitempty"`

	// 营业执照文件ID
	FileId *string `json:"file_id,omitempty"`

	// 是否涉及第三方权益 1. Yes: 是 2. No:
	IsInvolvedThird string `json:"is_involved_third"`

	// 授权委托书文件ID
	PowerAttorneyFileid *string `json:"power_attorney_fileid,omitempty"`

	// 签名名称
	SignatureName string `json:"signature_name"`

	// 签名来源。支持枚举值： 1. 0：企事业单位的全称或简称 2. 1：工信部备案网站的全称或简称 3. 2： APP应用的全称或简称 4. 3：公众号或小程序的全称或简称 5. 4：电商平台店铺名的全称或简称 6. 5：商标名的全称或简称
	SignatureSource int32 `json:"signature_source"`

	// 签名类型。支持枚举值： 1. VERIFY_CODE_TYPE: 验证码类 2. PROMOTION_TYPE: 推广类 3. NOTIFY_TYPE: 通知类
	SignatureType string `json:"signature_type"`

	// 签名来源标题内容
	SourceTitleContent *string `json:"source_title_content,omitempty"`
}

func (o SmsSignatureReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsSignatureReq struct{}"
	}

	return strings.Join([]string{"SmsSignatureReq", string(data)}, " ")
}
