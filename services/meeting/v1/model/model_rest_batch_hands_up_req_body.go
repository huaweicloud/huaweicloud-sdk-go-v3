package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestBatchHandsUpReqBody 批量举手请求。
type RestBatchHandsUpReqBody struct {

	// - 0: 放下手 - 1: 举手
	HandsState int32 `json:"handsState"`

	// 与会者标识列表。
	Pids []string `json:"pids"`
}

func (o RestBatchHandsUpReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestBatchHandsUpReqBody struct{}"
	}

	return strings.Join([]string{"RestBatchHandsUpReqBody", string(data)}, " ")
}
