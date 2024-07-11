package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHostClusterDetailResponse Response Object
type ShowHostClusterDetailResponse struct {

	// 请求成功失败状态
	Status *string `json:"status,omitempty"`

	Result         *HostClusterInfoDetail `json:"result,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowHostClusterDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHostClusterDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowHostClusterDetailResponse", string(data)}, " ")
}
