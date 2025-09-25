package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowActionIInfoResponse Response Object
type ShowActionIInfoResponse struct {
	Result *OctopusV3LogResponseBodyResult `json:"result,omitempty"`

	// 返回错误信息
	Error *string `json:"error,omitempty"`

	// 返回状态信息
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowActionIInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowActionIInfoResponse struct{}"
	}

	return strings.Join([]string{"ShowActionIInfoResponse", string(data)}, " ")
}
