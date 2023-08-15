package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmsSignatureResp struct {

	// 签名主键id，用于获取、修改、删除、申请激活签名的唯一标识
	Id *string `json:"id,omitempty"`

	// 创建时间[yyyy-MM-dd HH:mm:ss]
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间[yyyy-MM-dd HH:mm:ss]
	UpdateTime *string `json:"update_time,omitempty"`

	// 租户customer id
	CustomerId *string `json:"customer_id,omitempty"`

	Tenant *TenantBasic `json:"tenant,omitempty"`

	// 签名名称
	SignatureName *string `json:"signature_name,omitempty"`

	// 签名id
	SignatureId *string `json:"signature_id,omitempty"`

	// 签名类型
	SignatureType *string `json:"signature_type,omitempty"`

	// 应用名称
	AppName *string `json:"app_name,omitempty"`

	// 申请描述
	ApplyDesc *string `json:"apply_desc,omitempty"`

	// 国内短信通道号，仅当签名审核成功，运营人员配置完成后返回。
	ChannelNum *string `json:"channel_num,omitempty"`

	// 审核意见
	ReviewDesc *string `json:"review_desc,omitempty"`

	// 文件id
	FileId *string `json:"file_id,omitempty"`

	// 签名状态
	Status *string `json:"status,omitempty"`

	// 站点
	Site *string `json:"site,omitempty"`

	// 签名来源
	SignatureSource *int32 `json:"signature_source,omitempty"`

	// 是否涉及第三方权益
	IsInvolvedThird *string `json:"is_involved_third,omitempty"`

	// 授权委托书文件ID
	PowerAttorneyFileId *string `json:"power_attorney_file_id,omitempty"`

	// 催审状态
	UrgeStatus *string `json:"urge_status,omitempty"`

	// 催审时间
	UrgeTime *string `json:"urge_time,omitempty"`

	// 催审描述
	UrgeDesc *string `json:"urge_desc,omitempty"`

	// 应用key
	AppKey *string `json:"app_key,omitempty"`

	// 标题内容
	SourceTitleContent *string `json:"source_title_content,omitempty"`

	// 签名用途
	SignatureUsage *string `json:"signature_usage,omitempty"`
}

func (o SmsSignatureResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmsSignatureResp struct{}"
	}

	return strings.Join([]string{"SmsSignatureResp", string(data)}, " ")
}
