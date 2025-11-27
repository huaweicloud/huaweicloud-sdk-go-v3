package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchSubscribeReportRequestBody struct {

	// 是否订阅（true-订阅，false-取消订阅）
	Subscribe bool `json:"subscribe"`

	// 邮件模板ID列表
	EmailTemplateIdList []int32 `json:"email_template_id_list"`
}

func (o BatchSubscribeReportRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSubscribeReportRequestBody struct{}"
	}

	return strings.Join([]string{"BatchSubscribeReportRequestBody", string(data)}, " ")
}
