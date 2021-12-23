package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPublicipsResponse struct {
	// 本次请求的编号

	RequestId *string `json:"request_id,omitempty"`
	// 功能说明：弹性公网IP对象

	Publicips *[]PublicipSingleShowResp `json:"publicips,omitempty"`

	PageInfo *PageInfoOption `json:"page_info,omitempty"`
	// 公网IP总条目数

	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPublicipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicipsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicipsResponse", string(data)}, " ")
}
