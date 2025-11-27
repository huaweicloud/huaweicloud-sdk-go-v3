package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEmailRecordRequest Request Object
type ListEmailRecordRequest struct {

	// 数据库类型。支持MySQL、TaurusDB、GaussDB、MariaDB。
	DatastoreType string `json:"datastore_type"`

	// 开始时间（Unix timestamp），单位：毫秒。
	StartAt *int64 `json:"start_at,omitempty"`

	// 结束时间（Unix timestamp），单位：毫秒。
	EndAt *int64 `json:"end_at,omitempty"`

	// 发送状态
	SendStatus *int32 `json:"send_status,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListEmailRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEmailRecordRequest struct{}"
	}

	return strings.Join([]string{"ListEmailRecordRequest", string(data)}, " ")
}
