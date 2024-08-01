package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportTopRiskInstancesResponse Response Object
type ExportTopRiskInstancesResponse struct {

	// 风险实例列表。
	TopRiskInfo *[]TopRiskInfo `json:"top_risk_info,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportTopRiskInstancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportTopRiskInstancesResponse struct{}"
	}

	return strings.Join([]string{"ExportTopRiskInstancesResponse", string(data)}, " ")
}
