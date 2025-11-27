package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportStatsOpenResponse Response Object
type ExportStatsOpenResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ExportStatsOpenResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportStatsOpenResponse struct{}"
	}

	return strings.Join([]string{"ExportStatsOpenResponse", string(data)}, " ")
}
