package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetStatsConfigResponse Response Object
type SetStatsConfigResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o SetStatsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetStatsConfigResponse struct{}"
	}

	return strings.Join([]string{"SetStatsConfigResponse", string(data)}, " ")
}
