package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SaveOrderBody 客户服务单
type SaveOrderBody struct {

	// 服务单号，修改已保存的草稿时使用
	Number *string `json:"number,omitempty"`

	// 标题
	Title string `json:"title"`

	// 服务单类型编码
	ModelCode string `json:"model_code"`

	AccessWhitelist *AccessWhiteList `json:"access_whitelist,omitempty"`

	DeliveryInfo *DeliveryInfo `json:"delivery_info,omitempty"`

	// 操作对象
	OperationObjects *[]OperationObject `json:"operation_objects,omitempty"`

	// 操作内容附件
	OperationAttachments *[]UploadFileInfo `json:"operation_attachments,omitempty"`

	// 补充附件
	AdditionalAttachments *[]UploadFileInfo `json:"additional_attachments,omitempty"`

	// 补充说明
	AdditionalContent *string `json:"additional_content,omitempty"`

	// 手机国际码
	CountryCode *string `json:"country_code,omitempty"`

	// 联系电话
	Phone *string `json:"phone,omitempty"`

	// 扩展字段
	Fields *[]FieldEntity `json:"fields,omitempty"`

	// 是否同意服务条款
	IsAcceptTerms bool `json:"is_accept_terms"`
}

func (o SaveOrderBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SaveOrderBody struct{}"
	}

	return strings.Join([]string{"SaveOrderBody", string(data)}, " ")
}
