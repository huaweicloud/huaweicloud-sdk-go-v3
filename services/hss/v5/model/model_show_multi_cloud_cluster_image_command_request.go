package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMultiCloudClusterImageCommandRequest Request Object
type ShowMultiCloudClusterImageCommandRequest struct {

	// 镜像仓地址
	ImageRepo string `json:"image_repo"`

	// 组织名称
	Organization string `json:"organization"`

	// 用户名
	Username string `json:"username"`

	// 密码
	Password string `json:"password"`

	// **参数解释**: 插件类型 **约束限制**: 不涉及 **取值范围**: - docker: docker插件镜像 - agent: hostguard镜像 **默认取值**: 不涉及
	PlugType *string `json:"plug_type,omitempty"`
}

func (o ShowMultiCloudClusterImageCommandRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMultiCloudClusterImageCommandRequest struct{}"
	}

	return strings.Join([]string{"ShowMultiCloudClusterImageCommandRequest", string(data)}, " ")
}
