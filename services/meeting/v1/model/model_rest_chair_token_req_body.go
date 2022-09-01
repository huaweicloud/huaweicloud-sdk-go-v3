package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 申请释放主持人消息。
type RestChairTokenReqBody struct {

	// - 0: 释放主持人。 - 1: 申请主持人。
	ApplyChair int32 `json:"applyChair" xml:"applyChair"`

	// 当申请主持人时，携带主持人密码。
	ChairmanPwd *string `json:"chairmanPwd,omitempty" xml:"chairmanPwd"`
}

func (o RestChairTokenReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestChairTokenReqBody struct{}"
	}

	return strings.Join([]string{"RestChairTokenReqBody", string(data)}, " ")
}
