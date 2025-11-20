package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRuleImageResponseInfo 受配置检测影响的镜像信息
type CheckRuleImageResponseInfo struct {

	// **参数解释**: 组织名称 **取值范围**: 字符长度0-64
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像digest **取值范围**: 字符长度0-128
	ImageDigest *string `json:"image_digest,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-128
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本 **取值范围**: 字符长度0-128
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像类型 **取值范围**: - private_image：私有镜像 - shared_image：共享镜像 - instance_image：企业镜像 - harbor：Harbor镜像 - jfrog：Jfrog镜像 - cicd：cicd镜像
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 镜像仓库名称 **取值范围**: 字符长度0-128
	RegistryName *string `json:"registry_name,omitempty"`

	// **参数解释**: 版本大小 **取值范围**: 大小0-2147483547
	ImageSize *int32 `json:"image_size,omitempty"`
}

func (o CheckRuleImageResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRuleImageResponseInfo struct{}"
	}

	return strings.Join([]string{"CheckRuleImageResponseInfo", string(data)}, " ")
}
