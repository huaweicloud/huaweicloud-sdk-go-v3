package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProtectImageInfo **参数解释**: 防护的容器镜像信息 **约束限制**: 不涉及 **取值范围**: 不涉及 **默认取值**: 不涉及
type ProtectImageInfo struct {

	// **参数解释**： 镜像名称 **约束限制**: 不涉及 **取值范围**： 字符长度1-256位 **默认取值**: 不涉及
	ImageName string `json:"image_name"`

	// **参数解释**： 镜像版本 **约束限制**: 不涉及 **取值范围**： 字符长度0-64位 **默认取值**: 不涉及
	ImageVersion *string `json:"image_version,omitempty"`

	// **参数解释**： 镜像类型 **约束限制**: 不涉及 **取值范围**： - registry 仓库镜像 - local 本地镜像 - custom 自定义镜像  **默认取值**: custom
	ImageType *string `json:"image_type,omitempty"`

	// **参数解释**： 防护目录信息 **约束限制**: 不涉及 **取值范围**： 最少1条，最多50条 **默认取值**: 不涉及
	ProtectDirectoryInfoList []ProtectDirectoryInfo `json:"protect_directory_info_list"`

	// **参数解释**： 防护排除的文件类型列表 **约束限制**: 不涉及 **取值范围**： 最少0条，最多10条 **默认取值**: 不涉及
	ExcludeFileTypes *[]string `json:"exclude_file_types,omitempty"`
}

func (o ProtectImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProtectImageInfo struct{}"
	}

	return strings.Join([]string{"ProtectImageInfo", string(data)}, " ")
}
