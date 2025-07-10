package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListControlsForAccountResponse Response Object
type ListControlsForAccountResponse struct {

	// 治理策略概要。
	ControlSummaries *[]TargetControl `json:"control_summaries,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListControlsForAccountResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListControlsForAccountResponse struct{}"
	}

	return strings.Join([]string{"ListControlsForAccountResponse", string(data)}, " ")
}
