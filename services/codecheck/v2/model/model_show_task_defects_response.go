package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowTaskDefectsResponse struct {

	// 缺陷详情信息
	Defects *[]DefectInfoV2 `json:"defects,omitempty" xml:"defects"`

	// 总数
	Total          *int32 `json:"total,omitempty" xml:"total"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskDefectsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskDefectsResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskDefectsResponse", string(data)}, " ")
}
