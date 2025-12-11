package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateImagesInfo 多架构关联镜像信息
type AssociateImagesInfo struct {

	// **参数解释**: 多架构关联镜像的记录唯一标识ID **取值范围**: 最小值0，最大值9223372036854775807
	Id *int64 `json:"id,omitempty"`

	// **参数解释**: 多架构关联镜像的唯一标识ID **取值范围**: 字符长度0-64位
	ImageId *string `json:"image_id,omitempty"`

	// **参数解释**: 镜像名称 **取值范围**: 字符长度0-128位
	ImageName *string `json:"image_name,omitempty"`

	// **参数解释**: 镜像版本 **取值范围**: 字符长度0-64位
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**: 镜像类型 **取值范围**: - private_image：SWR私有镜像。 - shared_image：SWR共享镜像。 - instance_image：SWR企业版镜像。 - harbor：Harbor仓库镜像。 - jfrog：Jfrog镜像。
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**: 组织名称 **取值范围**: 字符长度0-64位
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**: 镜像digest **取值范围**: 字符长度0-128位
	ImageDigest *string `json:"image_digest,omitempty"`

	// **参数解释**： 扫描状态 **取值范围**： - unscan：未扫描。 - success：扫描完成。 - scanning：正在扫描。 - failed：扫描失败。 - download_failed：下载失败。 - image_oversized：镜像超大。 - waiting_for_scan：等待扫描。
	ScanStatus *string `json:"scan_status,omitempty"`

	// **参数解释**: 漏洞个数 **取值范围**: 取值0-2147483647
	VulNum *int32 `json:"vul_num,omitempty"`

	// **参数解释**: 基线扫描未通过数 **取值范围**: 取值0-2147483647
	UnsafeSettingNum *int32 `json:"unsafe_setting_num,omitempty"`

	// **参数解释**: 恶意文件数 **取值范围**: 取值0-2147483647
	MaliciousFileNum *int32 `json:"malicious_file_num,omitempty"`
}

func (o AssociateImagesInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateImagesInfo struct{}"
	}

	return strings.Join([]string{"AssociateImagesInfo", string(data)}, " ")
}
