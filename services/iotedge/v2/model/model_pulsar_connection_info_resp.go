package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PulsarConnectionInfoResp 外部推送通道返回详情
type PulsarConnectionInfoResp struct {

	// 鉴权token
	Token *string `json:"token,omitempty"`
}

func (o PulsarConnectionInfoResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PulsarConnectionInfoResp struct{}"
	}

	return strings.Join([]string{"PulsarConnectionInfoResp", string(data)}, " ")
}
