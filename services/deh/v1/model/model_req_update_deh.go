package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReqUpdateDeh 更新专属主机属性消息体。
type ReqUpdateDeh struct {
	DedicatedHost *ReqUpdateDehMessage `json:"dedicated_host"`
}

func (o ReqUpdateDeh) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqUpdateDeh struct{}"
	}

	return strings.Join([]string{"ReqUpdateDeh", string(data)}, " ")
}
