package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListHistoryOperationsResponse struct {

	// 参数修改历史的列表记录。
	Histories *[]ListHistoryOperationsResult `json:"histories,omitempty"`

	// 总记录数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListHistoryOperationsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryOperationsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryOperationsResponse", string(data)}, " ")
}
