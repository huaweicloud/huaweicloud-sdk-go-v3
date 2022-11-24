package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 举手请求。
type RestHandsUpReqBody struct {

	// - 0: 放下手 - 1: 举手
	HandsState int32 `json:"handsState"`
}

func (o RestHandsUpReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestHandsUpReqBody struct{}"
	}

	return strings.Join([]string{"RestHandsUpReqBody", string(data)}, " ")
}
