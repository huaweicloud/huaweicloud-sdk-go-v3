package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateCloudWafPostPaidResourceRequest struct {

	// 区域id
	Region string `json:"region"`

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreateCloudWafPostPaidResourceRequestbody `json:"body,omitempty"`
}

func (o CreateCloudWafPostPaidResourceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCloudWafPostPaidResourceRequest struct{}"
	}

	return strings.Join([]string{"CreateCloudWafPostPaidResourceRequest", string(data)}, " ")
}
