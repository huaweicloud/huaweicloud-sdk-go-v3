package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVolumeProductInfoResponse Response Object
type ListVolumeProductInfoResponse struct {

	// 磁盘产品信息列表
	Volumes        *[]VolumeProductInfo `json:"volumes,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListVolumeProductInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVolumeProductInfoResponse struct{}"
	}

	return strings.Join([]string{"ListVolumeProductInfoResponse", string(data)}, " ")
}
