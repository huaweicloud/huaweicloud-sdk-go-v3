package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChannelUpdateReq struct {

	// 通道描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o ChannelUpdateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelUpdateReq struct{}"
	}

	return strings.Join([]string{"ChannelUpdateReq", string(data)}, " ")
}
