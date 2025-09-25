package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCicdConfigurationCommandRequestBody 新增cicd配置
type CreateCicdConfigurationCommandRequestBody struct {

	// **参数解释**： 扫描类型 **约束限制**: 不涉及 **取值范围**: - local_image：本地镜像，镜像扫描参数 - remote_image：远程镜像仓，镜像扫描参数 - local_directory：本地目录，iac扫描参数 - remote_address：https远程地址 - git_repository_address：git仓库地址  **默认取值**: 不涉及
	ScanType *string `json:"scan_type,omitempty"`

	// **参数解释**： cicd标识 **约束限制**： 不涉及 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
	CicdId string `json:"cicd_id"`

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

	// **参数解释**: 是否进行镜像扫描 **约束限制**: 不涉及 **取值范围**:   - true：镜像扫描。   - false：不进行镜像扫描。  **默认取值**: 不涉及
	IsImageScan *bool `json:"is_image_scan,omitempty"`

	ImageScanInfo *ImageScanRequestInfo `json:"image_scan_info,omitempty"`

	// **参数解释**: 是否进行iac扫描 **约束限制**: 不涉及 **取值范围**:   - true：iac扫描。   - false：不进行iac扫描。  **默认取值**: 不涉及
	IsIacScan *bool `json:"is_iac_scan,omitempty"`

	IacScanInfo *IacScanRequestInfo `json:"iac_scan_info,omitempty"`
}

func (o CreateCicdConfigurationCommandRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCicdConfigurationCommandRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCicdConfigurationCommandRequestBody", string(data)}, " ")
}
