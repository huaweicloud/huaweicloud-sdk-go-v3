package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowOrderResponse Response Object
type ShowOrderResponse struct {

	// 服务单号
	Number *string `json:"number,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 工单类型:IDC运维 设备运维 设备检查 客户陪同
	Type *string `json:"type,omitempty"`

	// 具体操作类型:设备物理上下电
	SubType *string `json:"sub_type,omitempty"`

	// 工单类型编码
	ModelCode *string `json:"model_code,omitempty"`

	// 操作对象
	OperationObjects *[]OperationObject `json:"operation_objects,omitempty"`

	// 操作内容附件
	OperationAttachments *[]UploadFileInfo `json:"operation_attachments,omitempty"`

	AccessWhitelist *AccessWhiteList `json:"access_whitelist,omitempty"`

	DeliveryInfo *DeliveryInfo `json:"delivery_info,omitempty"`

	// 补充附件
	AdditionalAttachments *[]UploadFileInfo `json:"additional_attachments,omitempty"`

	// 补充说明
	AdditionalContent *string `json:"additional_content,omitempty"`

	// 手机国际码
	CountryCode *string `json:"country_code,omitempty"`

	// 联系电话
	Phone *string `json:"phone,omitempty"`

	// 当前阶段.DRAFT草稿、SUMMITED服务申请、IN_PROGRESS服务处理中、ACCEPTANCE服务验收、CLOSED服务关闭
	Stage *string `json:"stage,omitempty"`

	// 当前状态
	Status *string `json:"status,omitempty"`

	// 工单日志
	Logs *[]OrderLog `json:"logs,omitempty"`

	// 扩展字段
	Fields *[]FieldEntity `json:"fields,omitempty"`

	// 申请人
	Applicant *string `json:"applicant,omitempty"`

	// 是否同意服务条款
	IsAcceptTerms *bool `json:"is_accept_terms,omitempty"`

	// 申请时间/创建时间
	CreateTime *string `json:"create_time,omitempty"`

	Verification   *OrderVerification `json:"verification,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowOrderResponse struct{}"
	}

	return strings.Join([]string{"ShowOrderResponse", string(data)}, " ")
}
