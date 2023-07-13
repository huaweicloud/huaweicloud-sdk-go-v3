package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsEipsResponse Response Object
type ListDesktopsEipsResponse struct {

	// 桌面EIP。
	Eips *[]Eips `json:"eips,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopsEipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsEipsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopsEipsResponse", string(data)}, " ")
}
