package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRelatedProjectInfoResponse Response Object
type ListRelatedProjectInfoResponse struct {
	Result *ListRelatedProjectInfoResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRelatedProjectInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRelatedProjectInfoResponse struct{}"
	}

	return strings.Join([]string{"ListRelatedProjectInfoResponse", string(data)}, " ")
}
