package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrackerSmnChannelConfigBody SMN通道设置对象。跨帐号授予SMN主题发送通知的权限请参考《用户指南- 资源记录器- 开启/配置/修改资源记录器》中的“跨帐号授权”内容。
type TrackerSmnChannelConfigBody struct {

	// SMN主题的区域id
	RegionId string `json:"region_id"`

	// 创建或更新资源记录器用户的项目id
	ProjectId string `json:"project_id"`

	// SMN主题urn
	TopicUrn string `json:"topic_urn"`
}

func (o TrackerSmnChannelConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerSmnChannelConfigBody struct{}"
	}

	return strings.Join([]string{"TrackerSmnChannelConfigBody", string(data)}, " ")
}
