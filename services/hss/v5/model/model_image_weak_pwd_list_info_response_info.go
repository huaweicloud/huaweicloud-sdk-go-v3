package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageWeakPwdListInfoResponseInfo 弱口令的账号信息
type ImageWeakPwdListInfoResponseInfo struct {

	// **参数解释**: 组织名称（只有私有镜像和共享镜像有该字段，本地镜像没有） **取值范围**: 字符长度0-65535位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-65535位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本名称 **取值范围**: 字符长度0-128位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像id **取值范围**: 字符长度0-128位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 仓库镜像类型，包含如下: **取值范围**: - SwrPrivate : swr私有镜像。 - SwrShared : swr共享。 - SwrEnterprise : swr企业。 - Harbor : harbor仓库。 - Jfrog : jfrog仓库。 - Other : 其他仓库。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 最后一次检测时间，时间单位 毫秒（ms） **取值范围**: 最小值0，最大值2147483647
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 弱口令账号名称 **取值范围**: 字符长度0-32位
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: 账号类型 **取值范围**: - system : 系统。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**: 弱口令使用时长，单位天 **取值范围**: 最小值0，最大值2147483647
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释**: 脱敏弱口令 **取值范围**: 字符长度0-128位
	DesensitizedWeakPasswords *string `json:"desensitized_weak_passwords,omitempty"`

	// **参数解释**: 修改建议 **取值范围**: 字符长度0-65534位
	Suggestion *string `json:"suggestion,omitempty"`
}

func (o ImageWeakPwdListInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageWeakPwdListInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"ImageWeakPwdListInfoResponseInfo", string(data)}, " ")
}
