package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeletePoliciesRequest Request Object
type BatchDeletePoliciesRequest struct {

	// 您可以通过调用企业项目管理服务（EPS)的查询企业项目列表接口（ListEnterpriseProject）查询企业项目id
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *BatchDeletePoliciesRequestBody `json:"body,omitempty"`
}

func (o BatchDeletePoliciesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeletePoliciesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeletePoliciesRequest", string(data)}, " ")
}
