package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmsSignatureInfo 查询短信签名详情返回体。
type SmsSignatureInfo struct {

	// 签名名称。
	SignatureName *string `json:"signature_name,omitempty"`

	// 签名ID。
	SignatureId *string `json:"signature_id,omitempty"`

	// 签名类型。
	SignatureType *string `json:"signature_type,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 申请说明。
	ApplyDesc *string `json:"apply_desc,omitempty"`

	// 国内短信通道号，仅当签名审核成功，运营人员配置完成后返回。
	ChannelNum *string `json:"channel_num,omitempty"`

	// 营业执照文件ID。
	FileId *string `json:"file_id,omitempty"`

	// 状态。
	Status *SmsSignatureInfoStatus `json:"status,omitempty"`

	// 签名来源。
	SignatureSource *int32 `json:"signature_source,omitempty"`

	// 是否涉及第三方权益。
	IsInvolvedThird *string `json:"is_involved_third,omitempty"`

	// 授权委托书文件ID。
	PowerAttorneyFileId *string `json:"power_attorney_file_id,omitempty"`

	// 催审状态。
	UrgeStatus *string `json:"urge_status,omitempty"`

	// 催审时间。
	UrgeTime *string `json:"urge_time,omitempty"`

	// 催审描述。
	UrgeDesc *string `json:"urge_desc,omitempty"`

	// 审核意见。
	ReviewDesc *string `json:"review_desc,omitempty"`

	// 标题内容。
	SourceTitleContent *string `json:"source_title_content,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o SmsSignatureInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsSignatureInfo struct{}"
	}

	return strings.Join([]string{"SmsSignatureInfo", string(data)}, " ")
}

type SmsSignatureInfoStatus struct {
	value string
}

type SmsSignatureInfoStatusEnum struct {
	PENDING_REVIEW   SmsSignatureInfoStatus
	PROCESSING       SmsSignatureInfoStatus
	REVIEW_PASSED    SmsSignatureInfoStatus
	REVIEW_NOT_PASS  SmsSignatureInfoStatus
	TO_BE_ACTIVATED  SmsSignatureInfoStatus
	PENDING_ACTIVATE SmsSignatureInfoStatus
}

func GetSmsSignatureInfoStatusEnum() SmsSignatureInfoStatusEnum {
	return SmsSignatureInfoStatusEnum{
		PENDING_REVIEW: SmsSignatureInfoStatus{
			value: "PENDING_REVIEW：待审核",
		},
		PROCESSING: SmsSignatureInfoStatus{
			value: "PROCESSING：内容审核通过签名在处理中",
		},
		REVIEW_PASSED: SmsSignatureInfoStatus{
			value: "REVIEW_PASSED：处理完毕",
		},
		REVIEW_NOT_PASS: SmsSignatureInfoStatus{
			value: "REVIEW_NOT_PASS：审核不通过",
		},
		TO_BE_ACTIVATED: SmsSignatureInfoStatus{
			value: "TO_BE_ACTIVATED：待激活",
		},
		PENDING_ACTIVATE: SmsSignatureInfoStatus{
			value: "PENDING_ACTIVATE：激活审核中",
		},
	}
}

func (c SmsSignatureInfoStatus) Value() string {
	return c.value
}

func (c SmsSignatureInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmsSignatureInfoStatus) UnmarshalJSON(b []byte) error {
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
