package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudWafPostPaidResourceRequest Request Object
type DeleteCloudWafPostPaidResourceRequest struct {

	// 区域id
	Region *string `json:"region,omitempty"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o DeleteCloudWafPostPaidResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudWafPostPaidResourceRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudWafPostPaidResourceRequest", string(data)}, " ")
}
