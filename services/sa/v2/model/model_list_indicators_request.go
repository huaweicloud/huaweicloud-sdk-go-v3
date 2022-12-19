package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListIndicatorsRequest struct {

	// application/json;charset=UTF-8
	ContentType string `json:"content-type"`

	// workspace id
	WorkspaceId string `json:"workspace_id"`

	// sort order, ASC, DESC.
	Order *string `json:"order,omitempty"`

	// 起始时间
	FromDate *string `json:"from_date,omitempty"`

	// 结束时间
	ToDate *string `json:"to_date,omitempty"`

	Body *IndicatorListSearchRequest `json:"body,omitempty"`
}

func (o ListIndicatorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIndicatorsRequest struct{}"
	}

	return strings.Join([]string{"ListIndicatorsRequest", string(data)}, " ")
}
