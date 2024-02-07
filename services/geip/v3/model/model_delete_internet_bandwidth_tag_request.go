package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInternetBandwidthTagRequest Request Object
type DeleteInternetBandwidthTagRequest struct {

	// 全域公网带宽的id
	ResourceId string `json:"resource_id"`

	// 待删除标签的key
	TagKey string `json:"tag_key"`
}

func (o DeleteInternetBandwidthTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInternetBandwidthTagRequest struct{}"
	}

	return strings.Join([]string{"DeleteInternetBandwidthTagRequest", string(data)}, " ")
}
