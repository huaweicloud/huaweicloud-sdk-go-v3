package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrackerObsChannelConfigBody OBS设置对象。跨帐号授予OBS桶转储文件的权限请参考《用户指南- 资源记录器- 开启/配置/修改资源记录器》中的“跨帐号授权”内容。
type TrackerObsChannelConfigBody struct {

	// OBS桶名称
	BucketName string `json:"bucket_name"`

	// 区域id
	RegionId string `json:"region_id"`
}

func (o TrackerObsChannelConfigBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrackerObsChannelConfigBody struct{}"
	}

	return strings.Join([]string{"TrackerObsChannelConfigBody", string(data)}, " ")
}
