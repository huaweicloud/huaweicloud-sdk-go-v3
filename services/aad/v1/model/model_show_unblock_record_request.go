package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowUnblockRecordRequest Request Object
type ShowUnblockRecordRequest struct {

	// 租户id
	DomainId string `json:"domain_id"`

	// 开始时间
	StartTime int64 `json:"start_time"`

	// 结束时间
	EndTime int64 `json:"end_time"`
}

func (o ShowUnblockRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowUnblockRecordRequest struct{}"
	}

	return strings.Join([]string{"ShowUnblockRecordRequest", string(data)}, " ")
}
