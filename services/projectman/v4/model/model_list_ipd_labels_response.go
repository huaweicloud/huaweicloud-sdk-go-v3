package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIpdLabelsResponse Response Object
type ListIpdLabelsResponse struct {
	Result *LabelListResponse `json:"result,omitempty"`

	// 返回状态。
	Status *string `json:"status,omitempty"`

	// 请求失败时的错误信息。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListIpdLabelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIpdLabelsResponse struct{}"
	}

	return strings.Join([]string{"ListIpdLabelsResponse", string(data)}, " ")
}
