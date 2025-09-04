package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildDetailsResponse Response Object
type ShowBuildDetailsResponse struct {
	Result *JobStatusResultResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowBuildDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildDetailsResponse struct{}"
	}

	return strings.Join([]string{"ShowBuildDetailsResponse", string(data)}, " ")
}
