package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePrepaidCloudWafRequest Request Object
type CreatePrepaidCloudWafRequest struct {

	// 您可以通过调用企业项目管理服务（EPS）的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CreatePrepaidCloudWafRequestBody `json:"body,omitempty"`
}

func (o CreatePrepaidCloudWafRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePrepaidCloudWafRequest struct{}"
	}

	return strings.Join([]string{"CreatePrepaidCloudWafRequest", string(data)}, " ")
}
