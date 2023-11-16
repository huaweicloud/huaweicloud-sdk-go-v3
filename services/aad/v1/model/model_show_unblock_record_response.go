package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUnblockRecordResponse Response Object
type ShowUnblockRecordResponse struct {

	// 解封记录
	UnblockRecord *[]UnblockRecordResponseUnblockRecord `json:"unblock_record,omitempty"`

	// 总数
	Total *int32 `json:"total,omitempty"`

	// 租户id
	DomainId       *string `json:"domain_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowUnblockRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUnblockRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowUnblockRecordResponse", string(data)}, " ")
}
