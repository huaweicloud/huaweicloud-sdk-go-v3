package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SMN通道设置对象
type TrackerSmnChannelConfigBody struct {

	// 区域id
	RegionId string `json:"region_id" xml:"region_id"`

	// 项目id
	ProjectId string `json:"project_id" xml:"project_id"`

	// SMN主题urn
	TopicUrn string `json:"topic_urn" xml:"topic_urn"`
}

func (o TrackerSmnChannelConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerSmnChannelConfigBody struct{}"
	}

	return strings.Join([]string{"TrackerSmnChannelConfigBody", string(data)}, " ")
}
