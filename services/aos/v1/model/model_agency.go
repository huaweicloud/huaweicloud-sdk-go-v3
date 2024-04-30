package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Agency 委托授权的信息
type Agency struct {

	// 用户使用的provider的名字。如果用户给予的provider_name含有重复的值，则返回400
	ProviderName string `json:"provider_name"`

	// 对应provider所使用的IAM委托名称，资源编排服务会使用此委托的权限去访问、创建对应provider的资源。agency_name和agency_urn必须有且只有一个存在
	AgencyName *string `json:"agency_name,omitempty"`

	// 委托URN  当用户定义Agency时，agency_name和agency_urn 必须有且只有一个存在。  推荐用户在使用信任委托时给予agency_urn，agency_name只支持接收普通委托名称，如果给予了信任委托名称，则会在部署模板时失败。
	AgencyUrn *string `json:"agency_urn,omitempty"`
}

func (o Agency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Agency struct{}"
	}

	return strings.Join([]string{"Agency", string(data)}, " ")
}
