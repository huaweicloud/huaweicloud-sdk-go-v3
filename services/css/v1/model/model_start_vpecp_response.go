package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type StartVpecpResponse struct {
	// 操作行为。createVpcepservice表示已开启终端节点。

	Action         *string `json:"action,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartVpecpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartVpecpResponse struct{}"
	}

	return strings.Join([]string{"StartVpecpResponse", string(data)}, " ")
}
