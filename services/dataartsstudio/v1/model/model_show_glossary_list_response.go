package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGlossaryListResponse Response Object
type ShowGlossaryListResponse struct {

	// 总数
	Count *int32 `json:"count,omitempty"`

	// 分页参数limit
	Limit *int32 `json:"limit,omitempty"`

	// 分页参数offset
	Offset *int32 `json:"offset,omitempty"`

	// 指标配额
	Quota *int32 `json:"quota,omitempty"`

	// 标签信息列表
	Tags           *[]GlossaryInfo `json:"tags,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowGlossaryListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGlossaryListResponse struct{}"
	}

	return strings.Join([]string{"ShowGlossaryListResponse", string(data)}, " ")
}
