package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListOrchestrationsResponse Response Object
type ListOrchestrationsResponse struct {

	// 本次返回的列表长度
	Size int32 `json:"size"`

	// 满足条件的记录数
	Total int64 `json:"total"`

	// 本次查询到的编排规则列表。
	Orchestrations *[]OrchestrationBaseResp `json:"orchestrations,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListOrchestrationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOrchestrationsResponse struct{}"
	}

	return strings.Join([]string{"ListOrchestrationsResponse", string(data)}, " ")
}
