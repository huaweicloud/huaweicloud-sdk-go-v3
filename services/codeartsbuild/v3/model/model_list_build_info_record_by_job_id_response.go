package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuildInfoRecordByJobIdResponse Response Object
type ListBuildInfoRecordByJobIdResponse struct {
	Result *ListBuildInfoRecordBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBuildInfoRecordByJobIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuildInfoRecordByJobIdResponse struct{}"
	}

	return strings.Join([]string{"ListBuildInfoRecordByJobIdResponse", string(data)}, " ")
}
