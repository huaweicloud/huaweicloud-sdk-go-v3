package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductVersionResp **参数解释**： 规格版本信息。 **取值范围**： 不涉及。
type ProductVersionResp struct {

	// **参数解释**： 产品规格在该版本支持的最小CN数量。 **取值范围**： 不涉及。
	MinCn *int32 `json:"min_cn,omitempty"`

	// **参数解释**： 产品规格在该版本支持的最大CN数量。 **取值范围**： 不涉及。
	MaxCn *int32 `json:"max_cn,omitempty"`

	// **参数解释**： 产品规格该版本支持的类型。 **取值范围**： 1：稳定版； 0：最新版本。
	VersionType *string `json:"version_type,omitempty"`

	// **参数解释**： 产品规格版本号名称。 **取值范围**： 不涉及。
	DatastoreVersion *string `json:"datastore_version,omitempty"`
}

func (o ProductVersionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductVersionResp struct{}"
	}

	return strings.Join([]string{"ProductVersionResp", string(data)}, " ")
}
