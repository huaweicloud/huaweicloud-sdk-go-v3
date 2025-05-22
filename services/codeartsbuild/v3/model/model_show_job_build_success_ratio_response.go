package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobBuildSuccessRatioResponse Response Object
type ShowJobBuildSuccessRatioResponse struct {
	Result *ShowJobBuildSuccessRatioResult `json:"result,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobBuildSuccessRatioResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildSuccessRatioResponse struct{}"
	}

	return strings.Join([]string{"ShowJobBuildSuccessRatioResponse", string(data)}, " ")
}
