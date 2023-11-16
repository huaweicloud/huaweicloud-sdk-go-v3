package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Bw 黑白名单详情
type Bw struct {

	// 黑名单列表
	BlackIpList []string `json:"black_ip_list"`

	// 白名单列表
	WhiteIpList []string `json:"white_ip_list"`
}

func (o Bw) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Bw struct{}"
	}

	return strings.Join([]string{"Bw", string(data)}, " ")
}
