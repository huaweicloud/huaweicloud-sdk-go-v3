package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OBS设置对象
type TrackerObsChannelConfigBody struct {

	// OBS桶名称
	BucketName string `json:"bucket_name"`

	// region id
	RegionId string `json:"region_id"`
}

func (o TrackerObsChannelConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerObsChannelConfigBody struct{}"
	}

	return strings.Join([]string{"TrackerObsChannelConfigBody", string(data)}, " ")
}
