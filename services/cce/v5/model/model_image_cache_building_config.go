package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageCacheBuildingConfig 镜像缓存构建配置信息。
type ImageCacheBuildingConfig struct {

	// **参数解释：** 构建镜像缓存所启动的临时Pod所在的Autopilot集群的UID。 **约束限制：** 要求使用的Autopilot集群版本为v1.28.8-r0或v1.31.4-r0以上版本。 **取值范围：** 不涉及 **默认取值：** 不涉及
	Cluster string `json:"cluster"`

	// 下载所需缓存镜像的访问凭证列表，不填写或无有效凭证时仅支持下载公共镜像。
	ImagePullSecrets *[]string `json:"image_pull_secrets,omitempty"`
}

func (o ImageCacheBuildingConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageCacheBuildingConfig struct{}"
	}

	return strings.Join([]string{"ImageCacheBuildingConfig", string(data)}, " ")
}
