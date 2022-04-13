package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 批量挂断会场消息体。
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
