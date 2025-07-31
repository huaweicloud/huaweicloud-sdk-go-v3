package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCceClusterDetectRiskResponse Response Object
type ListCceClusterDetectRiskResponse struct {

	// 集群总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 集群信息列表
	DataList       *[]NodeDetectRiskResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListCceClusterDetectRiskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCceClusterDetectRiskResponse struct{}"
	}

	return strings.Join([]string{"ListCceClusterDetectRiskResponse", string(data)}, " ")
}
