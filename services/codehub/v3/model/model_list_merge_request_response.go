package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMergeRequestResponse Response Object
type ListMergeRequestResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *MergeResult `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMergeRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMergeRequestResponse struct{}"
	}

	return strings.Join([]string{"ListMergeRequestResponse", string(data)}, " ")
}
