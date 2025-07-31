package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionItem struct {

	// **参数解释：** 版本号。 **取值范围：** v2。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 版本状态。 **取值范围：** - CURRENT：表示该版本为主推版本。 - SUPPORTED：表示为老版本，但是现在还继续支持。 - DEPRECATED：表示为废弃版本，存在后续删除的可能。
	Status *string `json:"status,omitempty"`

	// **参数解释：** API的URL地址。 **取值范围：** 不涉及。
	Links *[]LinksItem `json:"links,omitempty"`

	// **参数解释：** 版本发布时间。 **取值范围：** 不涉及。
	Updated *string `json:"updated,omitempty"`

	// **参数解释：** 支持的最大微版本号。 **取值范围：** 若该版本API不支持微版本，则为空。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 支持的最小微版本号。 **取值范围：** 若该版本API不支持微版本，则为空。
	MinVersion *string `json:"min_version,omitempty"`
}

func (o VersionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionItem struct{}"
	}

	return strings.Join([]string{"VersionItem", string(data)}, " ")
}
