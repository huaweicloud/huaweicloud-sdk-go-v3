package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetScriptJobStatisticsResponse Response Object
type GetScriptJobStatisticsResponse struct {
	Data           *JobScriptOrderStatisticsModel `json:"data,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o GetScriptJobStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetScriptJobStatisticsResponse struct{}"
	}

	return strings.Join([]string{"GetScriptJobStatisticsResponse", string(data)}, " ")
}
