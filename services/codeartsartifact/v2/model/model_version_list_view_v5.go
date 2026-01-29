package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VersionListViewV5 struct {

	// **参数解释**： 发布库版本的状态。 **取值范围**： - test：测试包。 - prod：发布包。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 发布库版本的名称。 **取值范围**： 不涉及。
	BuildVersion *string `json:"build_version,omitempty"`

	// **参数解释**： 版本下的文件个数。 **取值范围**： 不涉及。
	FilesCount *int32 `json:"files_count,omitempty"`
}

func (o VersionListViewV5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionListViewV5 struct{}"
	}

	return strings.Join([]string{"VersionListViewV5", string(data)}, " ")
}
