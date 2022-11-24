package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowMergeRequestResponse struct {
	Error *Error `json:"error,omitempty"`

	Result *MergeInfoResult `json:"result,omitempty"`

	// 响应状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMergeRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMergeRequestResponse struct{}"
	}

	return strings.Join([]string{"ShowMergeRequestResponse", string(data)}, " ")
}
