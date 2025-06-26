package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClustersRequest Request Object
type ListClustersRequest struct {

	// **参数解释**： 企业项目ID。查询所有绑定企业项目的集群，则值为all_granted_eps。 **约束限制**： 不涉及。 **取值范围**： all_granted_eps：所有企业项目。 0：表示默认企业项目“default”的ID。 其它：过滤对应企业项目下的数据。 **默认取值**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
}

func (o ListClustersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClustersRequest struct{}"
	}

	return strings.Join([]string{"ListClustersRequest", string(data)}, " ")
}
