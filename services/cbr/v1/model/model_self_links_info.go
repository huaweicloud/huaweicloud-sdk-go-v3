package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 连接信息
type SelfLinksInfo struct {

	// 连接地址
	Self *string `json:"self,omitempty"`
}

func (o SelfLinksInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelfLinksInfo struct{}"
	}

	return strings.Join([]string{"SelfLinksInfo", string(data)}, " ")
}
