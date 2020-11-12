/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// SMN通道设置对象
type TrackerSmnChannelConfigBody struct {
	// region id
	RegionId string `json:"region_id"`
	// project id
	ProjectId string `json:"project_id"`
	// SMN 主题urn
	TopicUrn string `json:"topic_urn"`
}

func (o TrackerSmnChannelConfigBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"TrackerSmnChannelConfigBody", string(data)}, " ")
}
