package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImagePwdComplexityInfoResponseInfo 镜像的口令复杂度策略
type ImagePwdComplexityInfoResponseInfo struct {

	// **参数解释**: 镜像ID **取值范围**: 字符长度0-128位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 组织名称（只有私有镜像和共享镜像有该字段，本地镜像没有） **取值范围**: 字符长度0-65535位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-65535位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本名称 **取值范围**: 字符长度0-256位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 仓库镜像类型 **取值范围**: - SwrPrivate : swr私有镜像。 - SwrShared : swr共享。 - SwrEnterprise : swr企业。 - Harbor : harbor仓库。 - Jfrog : jfrog仓库。 - Other : 其他仓库。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 最后一次检测时间，时间单位 毫秒（ms） **取值范围**: 最小值0，最大值2147483647
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**： 口令最小长度 **取值范围**： - true：是。 - false：否。
	MinLength *bool `json:"min_length,omitempty"`

	// **参数解释**： 大写字母 **取值范围**： - true：是。 - false：否。
	UppercaseLetter *bool `json:"uppercase_letter,omitempty"`

	// **参数解释**： 小写字母 **取值范围**： - true：是。 - false：否。
	LowercaseLetter *bool `json:"lowercase_letter,omitempty"`

	// **参数解释**： 数字 **取值范围**： - true：是。 - false：否。
	Number *bool `json:"number,omitempty"`

	// **参数解释**： 特殊字符 **取值范围**： - true：是。 - false：否。
	SpecialCharacter *bool `json:"special_character,omitempty"`

	// **参数解释**: 修改建议 **取值范围**: 字符长度0-65534位
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o ImagePwdComplexityInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImagePwdComplexityInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"ImagePwdComplexityInfoResponseInfo", string(data)}, " ")
}
