package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopOperationsResponse Response Object
type ListDesktopOperationsResponse struct {

	// 操作记录。
	Operations *[]OperationForList `json:"operations,omitempty"`

	// 总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListDesktopOperationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopOperationsResponse struct{}"
	}

	return strings.Join([]string{"ListDesktopOperationsResponse", string(data)}, " ")
}
