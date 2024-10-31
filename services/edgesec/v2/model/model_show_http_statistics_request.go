package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpStatisticsRequest Request Object
type ShowHttpStatisticsRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id，默认为0
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 起始时间
	From int64 `json:"from"`

	// 结束时间
	To int64 `json:"to"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`
}

func (o ShowHttpStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpStatisticsRequest", string(data)}, " ")
}
