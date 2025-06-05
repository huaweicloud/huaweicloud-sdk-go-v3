package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildInfoRecordResponse Response Object
type ShowBuildInfoRecordResponse struct {
	Result *BuildInfoRecord `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBuildInfoRecordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildInfoRecordResponse struct{}"
	}

	return strings.Join([]string{"ShowBuildInfoRecordResponse", string(data)}, " ")
}
