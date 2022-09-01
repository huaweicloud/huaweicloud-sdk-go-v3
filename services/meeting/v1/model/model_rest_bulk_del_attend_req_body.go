package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 删除会场消息体。
type RestBulkDelAttendReqBody struct {

	// 待删除列表
	BulkDelAttendInfo []DelAttendInfo `json:"bulkDelAttendInfo" xml:"bulkDelAttendInfo"`
}

func (o RestBulkDelAttendReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestBulkDelAttendReqBody struct{}"
	}

	return strings.Join([]string{"RestBulkDelAttendReqBody", string(data)}, " ")
}
