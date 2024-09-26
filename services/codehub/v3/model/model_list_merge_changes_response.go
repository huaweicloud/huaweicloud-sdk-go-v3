package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeChangesResponse Response Object
type ListMergeChangesResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *ResponseMergeRequestChanges `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMergeChangesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeChangesResponse struct{}"
	}

	return strings.Join([]string{"ListMergeChangesResponse", string(data)}, " ")
}
