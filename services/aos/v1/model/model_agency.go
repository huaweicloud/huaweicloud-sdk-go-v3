package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Agency 委托授权的信息
type Agency struct {

	// 用户使用的provider的名字。如果用户给与的provider_name含有重复的值，则返回400
	ProviderName string `json:"provider_name"`

	// 对应provider所使用的IAM委托名称，资源编排服务会使用此委托的权限去访问、创建对应provider的资源
	AgencyName string `json:"agency_name"`
}

func (o Agency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agency struct{}"
	}

	return strings.Join([]string{"Agency", string(data)}, " ")
}
