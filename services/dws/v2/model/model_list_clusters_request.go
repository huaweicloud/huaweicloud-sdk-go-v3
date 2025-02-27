package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersRequest Request Object
type ListClustersRequest struct {

	// 企业项目ID，查询所有绑定eps集群，值为all_granted_eps
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
