package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterRiskAffectResourcesResponse Response Object
type ListClusterRiskAffectResourcesResponse struct {

	// 集群风险影响的资源列表数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 集群风险影响的资源列表
	DataList       *[]ListClusterRiskAffectResourcesResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                                                   `json:"-"`
}

func (o ListClusterRiskAffectResourcesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterRiskAffectResourcesResponse struct{}"
	}

	return strings.Join([]string{"ListClusterRiskAffectResourcesResponse", string(data)}, " ")
}
