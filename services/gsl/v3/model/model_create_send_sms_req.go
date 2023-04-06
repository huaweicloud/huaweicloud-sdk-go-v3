package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateSendSmsReq struct {

	// 模板id
	TemplateId *int64 `json:"template_id,omitempty"`

	// 短信内容
	SmsContent string `json:"sms_content"`

	// 容器ID
	Cids *[]string `json:"cids,omitempty"`

	// 批次号
	OrderId *int64 `json:"order_id,omitempty"`

	// 临时文件ID
	FileTempId *int64 `json:"file_temp_id,omitempty"`
}

func (o CreateSendSmsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSendSmsReq struct{}"
	}

	return strings.Join([]string{"CreateSendSmsReq", string(data)}, " ")
}
