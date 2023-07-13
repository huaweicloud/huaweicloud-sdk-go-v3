package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestChairTokenReqBody 申请主持人请求。
type RestChairTokenReqBody struct {

	// - 0: 释放主持人 - 1: 申请主持人
	ApplyChair int32 `json:"applyChair"`

	// 当申请主持人时，携带主持人密码。
	ChairmanPwd *string `json:"chairmanPwd,omitempty"`
}

func (o RestChairTokenReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestChairTokenReqBody struct{}"
	}

	return strings.Join([]string{"RestChairTokenReqBody", string(data)}, " ")
}
