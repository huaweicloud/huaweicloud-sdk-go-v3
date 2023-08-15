package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSummaryResponse Response Object
type ShowSummaryResponse struct {

	// 总容量大小
	Size *int32 `json:"size,omitempty"`

	// 总使用量
	UsedSize       *int32 `json:"used_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowSummaryResponse", string(data)}, " ")
}
