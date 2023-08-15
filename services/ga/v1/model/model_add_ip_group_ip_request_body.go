package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddIpGroupIpRequestBody add ip group ip request
type AddIpGroupIpRequestBody struct {

	// IP地址组中的IP网段列表，一次最多支持添加20个条目。
	IpList []CreateIpGroupIpOption `json:"ip_list"`
}

func (o AddIpGroupIpRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddIpGroupIpRequestBody struct{}"
	}

	return strings.Join([]string{"AddIpGroupIpRequestBody", string(data)}, " ")
}
