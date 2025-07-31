package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterRisksResponse Response Object
type ListClusterRisksResponse struct {

	// 集群风险列表数量
	TotalNum *int32 `json:"total_num,omitempty"`

	// 集群风险列表
	DataList       *[]ListClusterRisksResponseInfoDataList `json:"data_list,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListClusterRisksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterRisksResponse struct{}"
	}

	return strings.Join([]string{"ListClusterRisksResponse", string(data)}, " ")
}
