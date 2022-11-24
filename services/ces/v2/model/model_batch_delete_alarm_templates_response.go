package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchDeleteAlarmTemplatesResponse struct {

	// 成功删除的告警模板ID列表
	TemplateIds    *[]string `json:"template_ids,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteAlarmTemplatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAlarmTemplatesResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteAlarmTemplatesResponse", string(data)}, " ")
}
