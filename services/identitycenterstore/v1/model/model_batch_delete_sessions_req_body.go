package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteSessionsReqBody struct {

	// 用户登录会话标识符（ID）
	SessionIds []string `json:"session_ids"`
}

func (o BatchDeleteSessionsReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteSessionsReqBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteSessionsReqBody", string(data)}, " ")
}
