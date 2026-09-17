package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSpaceTrendResponse Response Object
type ShowSpaceTrendResponse struct {

	// 空间趋势指标列表
	Series         *[]SpaceTrend `json:"series,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowSpaceTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSpaceTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowSpaceTrendResponse", string(data)}, " ")
}
