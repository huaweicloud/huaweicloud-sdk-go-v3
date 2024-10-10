package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IpBindingV3Body ip绑定请求体
type IpBindingV3Body struct {

	// 防护ip列表
	IpList []string `json:"ip_list"`
}

func (o IpBindingV3Body) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IpBindingV3Body struct{}"
	}

	return strings.Join([]string{"IpBindingV3Body", string(data)}, " ")
}
