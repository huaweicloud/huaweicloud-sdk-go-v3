package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImageAppResponseInfo 软件信息
type ImageAppResponseInfo struct {

	// **参数解释**: 软件名称 **取值范围**: 字符长度0-128位
	AppName *string `json:"app_name,omitempty"`

	// **参数解释**: 软件类型 **取值范围**: 字符长度0-128位
	AppType *string `json:"app_type,omitempty"`

	// **参数解释**: 软件版本 **取值范围**: 字符长度0-128位
	AppVersion *string `json:"app_version,omitempty"`

	// **参数解释**: 漏洞总数 **取值范围**: 最小值0，最大值20000
	VulNum *int32 `json:"vul_num,omitempty"`

	// **参数解释**: 软件路径 **取值范围**: 字符长度0-128位
	AppPath *string `json:"app_path,omitempty"`

	// **参数解释**: 层digest **取值范围**: 字符长度0-128位
	LayerDigest *string `json:"layer_digest,omitempty"`

	// **参数解释**: 是否合规，false为不合规 **取值范围**: - true：是。 - false：否。
	IsCompliant *bool `json:"is_compliant,omitempty"`

	// **参数解释**: 最后一次检测时间，时间单位 毫秒（ms） **取值范围**: 最小值0，最大值9223372036854775807
	LatestScanTime *int64 `json:"latest_scan_time,omitempty"`

	// **参数解释**: 仓库镜像类型，包含如下: **取值范围**: - SwrPrivate : swr私有镜像。 - SwrShared : swr共享。 - SwrEnterprise : swr企业。 - Harbor : harbor仓库。 - Jfrog : jfrog仓库。 - Other : 其他仓库。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 组织名称 **取值范围**: 字符长度0-65535位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-65535位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本名称 **取值范围**: 字符长度0-65535位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像id **取值范围**: 字符长度0-65535位
	ImageId *string `json:"image_id,omitempty"`
}

func (o ImageAppResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImageAppResponseInfo struct{}"
	}

	return strings.Join([]string{"ImageAppResponseInfo", string(data)}, " ")
}
