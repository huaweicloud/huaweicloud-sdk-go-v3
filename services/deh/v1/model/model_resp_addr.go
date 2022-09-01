package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 云服务器的vpc信息。
type RespAddr struct {

	// 云服务器的vpc ip。
	Addr string `json:"addr" xml:"addr"`

	// 云服务器的vpc版本。
	Version string `json:"version" xml:"version"`
}

func (o RespAddr) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RespAddr struct{}"
	}

	return strings.Join([]string{"RespAddr", string(data)}, " ")
}
