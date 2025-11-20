package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageScanRequestInfo 镜像扫描信息
type ImageScanRequestInfo struct {

	// **参数解释**: 扫描类型 **约束限制**: 不涉及 **取值范围**: - local_image：本地镜像，镜像扫描参数。 - remote_image：远程镜像仓，镜像扫描参数。  **默认取值**: 不涉及
	ScanType *string `json:"scan_type,omitempty"`

	// **参数解释**: 镜像名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 流水线是否阻断 **约束限制**: 不涉及 **取值范围**:   - true：流水线被阻断。   - false：流水线未阻断。  **默认取值**: 不涉及
	IsBlocking *bool `json:"is_blocking,omitempty"`

	// **参数解释**: 镜像仓地址 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	RepositoryAddress *string `json:"repository_address,omitempty"`

	// **参数解释**: 镜像仓登录用户名 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	LoginUserName *string `json:"login_user_name,omitempty"`

	// **参数解释**: 镜像仓登录密码 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	LoginPassword *string `json:"login_password,omitempty"`

	// **参数解释**: 组织名称 **约束限制**: 不涉及 **取值范围**: 字符长度1-128位 **默认取值**: 不涉及
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 流水线类型 **约束限制**: 不涉及 **取值范围**: - jenkins：Jenkins流水线。 - codearts：CodeArts流水线。  **默认取值**: 不涉及
	PipelineType *string `json:"pipeline_type,omitempty"`
}

func (o ImageScanRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageScanRequestInfo struct{}"
	}

	return strings.Join([]string{"ImageScanRequestInfo", string(data)}, " ")
}
