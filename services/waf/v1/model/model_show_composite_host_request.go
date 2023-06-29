package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCompositeHostRequest Request Object
type ShowCompositeHostRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 域名id，通过查询云模式防护域名列表（ListHost）获取域名id或者通过独享模式域名列表（ListPremiumHost）获取域名id
	HostId string `json:"host_id"`
}

func (o ShowCompositeHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCompositeHostRequest struct{}"
	}

	return strings.Join([]string{"ShowCompositeHostRequest", string(data)}, " ")
}
