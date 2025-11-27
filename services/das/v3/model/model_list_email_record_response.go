package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEmailRecordResponse Response Object
type ListEmailRecordResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 邮件推送记录列表
	EmailRecordList *[]EmailRecord `json:"email_record_list,omitempty"`
	HttpStatusCode  int            `json:"-"`
}

func (o ListEmailRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEmailRecordResponse struct{}"
	}

	return strings.Join([]string{"ListEmailRecordResponse", string(data)}, " ")
}
