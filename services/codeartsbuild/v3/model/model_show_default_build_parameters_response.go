package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDefaultBuildParametersResponse Response Object
type ShowDefaultBuildParametersResponse struct {

	// 参数
	Result *[]ShowDefaultBuildParametersResult `json:"result,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDefaultBuildParametersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDefaultBuildParametersResponse struct{}"
	}

	return strings.Join([]string{"ShowDefaultBuildParametersResponse", string(data)}, " ")
}
