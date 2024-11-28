package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceClusterConfigRequest Request Object
type ListCceClusterConfigRequest struct {

	// Region ID
	Region string `json:"region"`

	// 企业项目ID，查询所有企业项目时填写：all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	Body *CceClusterInfoListRequestBody `json:"body,omitempty"`
}

func (o ListCceClusterConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceClusterConfigRequest struct{}"
	}

	return strings.Join([]string{"ListCceClusterConfigRequest", string(data)}, " ")
}
