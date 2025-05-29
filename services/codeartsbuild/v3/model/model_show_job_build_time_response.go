package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowJobBuildTimeResponse Response Object
type ShowJobBuildTimeResponse struct {
	Result *ShowJobBuildTimeResult `json:"result,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowJobBuildTimeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowJobBuildTimeResponse struct{}"
	}

	return strings.Join([]string{"ShowJobBuildTimeResponse", string(data)}, " ")
}
