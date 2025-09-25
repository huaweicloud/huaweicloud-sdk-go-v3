package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSecurityCheckClusterReportsResponse Response Object
type ListSecurityCheckClusterReportsResponse struct {

	// **参数解释**： 集群总数 **取值范围**： 不涉及
	TotalNum *int64 `json:"total_num,omitempty"`

	// **参数解释**： 安全体检报告列表 **取值范围**： 不涉及
	DataList       *[]SecurityCheckClusterReports `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListSecurityCheckClusterReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSecurityCheckClusterReportsResponse struct{}"
	}

	return strings.Join([]string{"ListSecurityCheckClusterReportsResponse", string(data)}, " ")
}
