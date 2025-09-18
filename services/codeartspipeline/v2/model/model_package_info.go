package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PackageInfo 流水线产物
type PackageInfo struct {

	// **参数解释**： 产物名。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 产物类型。 **取值范围**： 不涉及。
	PackageType *string `json:"packageType,omitempty"`

	// **参数解释**： 产物版本号。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 产物下载地址。 **取值范围**： 不涉及。
	DownloadUrl *string `json:"downloadUrl,omitempty"`
}

func (o PackageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PackageInfo struct{}"
	}

	return strings.Join([]string{"PackageInfo", string(data)}, " ")
}
