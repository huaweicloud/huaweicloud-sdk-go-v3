package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBuildParamsListResponse Response Object
type ShowBuildParamsListResponse struct {

	// 状态
	Status *string `json:"status,omitempty"`

	// 错误
	Error *interface{} `json:"error,omitempty"`

	Result         *BuildParamsBodyResult `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowBuildParamsListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBuildParamsListResponse struct{}"
	}

	return strings.Join([]string{"ShowBuildParamsListResponse", string(data)}, " ")
}
