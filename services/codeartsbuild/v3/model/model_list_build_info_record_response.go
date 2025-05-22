package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBuildInfoRecordResponse Response Object
type ListBuildInfoRecordResponse struct {
	Result *ListBuildInfoRecordBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListBuildInfoRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBuildInfoRecordResponse struct{}"
	}

	return strings.Join([]string{"ListBuildInfoRecordResponse", string(data)}, " ")
}
