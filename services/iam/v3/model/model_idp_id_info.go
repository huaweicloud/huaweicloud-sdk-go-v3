package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 身份提供商信息。
type IdpIdInfo struct {
	// 身份提供商id。

	Id string `json:"id"`
}

func (o IdpIdInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpIdInfo struct{}"
	}

	return strings.Join([]string{"IdpIdInfo", string(data)}, " ")
}
