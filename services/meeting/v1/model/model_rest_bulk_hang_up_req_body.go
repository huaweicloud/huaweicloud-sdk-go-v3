package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 邀请与会者请求。
type RestBulkHangUpReqBody struct {

	// 批量挂断会场列表，列表元素为与会者标识。
	BulkHangUpParticipants []string `json:"bulkHangUpParticipants"`
}

func (o RestBulkHangUpReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestBulkHangUpReqBody struct{}"
	}

	return strings.Join([]string{"RestBulkHangUpReqBody", string(data)}, " ")
}
