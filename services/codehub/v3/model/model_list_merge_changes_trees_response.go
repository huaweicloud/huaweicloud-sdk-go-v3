package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeChangesTreesResponse Response Object
type ListMergeChangesTreesResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *MergeChangesTreesDto `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMergeChangesTreesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeChangesTreesResponse struct{}"
	}

	return strings.Join([]string{"ListMergeChangesTreesResponse", string(data)}, " ")
}
