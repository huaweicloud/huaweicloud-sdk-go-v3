package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetStatsConfigRequest Request Object
type SetStatsConfigRequest struct {
	Body *SetStatsConfigBody `json:"body,omitempty"`
}

func (o SetStatsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetStatsConfigRequest struct{}"
	}

	return strings.Join([]string{"SetStatsConfigRequest", string(data)}, " ")
}
