package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCheckDesktopRejoinDomainResponse Response Object
type BatchCheckDesktopRejoinDomainResponse struct {

	// 检查总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 检查结果列表。
	Desktops       *[]BatchCheckRejoinDomainResult `json:"desktops,omitempty"`
	HttpStatusCode int                             `json:"-"`
}

func (o BatchCheckDesktopRejoinDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCheckDesktopRejoinDomainResponse struct{}"
	}

	return strings.Join([]string{"BatchCheckDesktopRejoinDomainResponse", string(data)}, " ")
}
