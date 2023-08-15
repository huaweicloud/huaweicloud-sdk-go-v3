package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RemoveIpGroupIpRequestBody remove ip group ip request
type RemoveIpGroupIpRequestBody struct {

	// IP地址组中的IP网段列表，一次最多支持删除20个条目。
	IpList []string `json:"ip_list"`
}

func (o RemoveIpGroupIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RemoveIpGroupIpRequestBody struct{}"
	}

	return strings.Join([]string{"RemoveIpGroupIpRequestBody", string(data)}, " ")
}
