package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchListEdgeAppVersionsResponse struct {
	// 总记录数

	Count *int32 `json:"count,omitempty"`

	PageInfo *PageInfoDto `json:"page_info,omitempty"`
	// 每页记录数

	Versions       *[]QueryEdgeAppVersionBriefResponseDto `json:"versions,omitempty"`
	HttpStatusCode int                                    `json:"-"`
}

func (o BatchListEdgeAppVersionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchListEdgeAppVersionsResponse struct{}"
	}

	return strings.Join([]string{"BatchListEdgeAppVersionsResponse", string(data)}, " ")
}
