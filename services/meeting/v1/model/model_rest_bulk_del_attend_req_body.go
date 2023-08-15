package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestBulkDelAttendReqBody 删除与会者请求。
type RestBulkDelAttendReqBody struct {

	// 待删除会场列表。
	BulkDelAttendInfo []DelAttendInfo `json:"bulkDelAttendInfo"`
}

func (o RestBulkDelAttendReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestBulkDelAttendReqBody struct{}"
	}

	return strings.Join([]string{"RestBulkDelAttendReqBody", string(data)}, " ")
}
