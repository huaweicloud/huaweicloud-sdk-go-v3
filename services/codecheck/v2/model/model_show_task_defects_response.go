package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskDefectsResponse struct {
	// 缺陷详情信息

	Defects *[]DefectInfoV2 `json:"defects,omitempty"`
	// 总数

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskDefectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDefectsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDefectsResponse", string(data)}, " ")
}
