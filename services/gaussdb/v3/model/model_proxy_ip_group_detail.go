package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ProxyIpGroupDetail struct {

	// ipGroup的id
	Id string `json:"id"`

	// ipGroup的名称
	Name string `json:"name"`

	// ipGroup内部的ip列表
	IpList []IpGroupItem `json:"ip_list"`
}

func (o ProxyIpGroupDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProxyIpGroupDetail struct{}"
	}

	return strings.Join([]string{"ProxyIpGroupDetail", string(data)}, " ")
}
