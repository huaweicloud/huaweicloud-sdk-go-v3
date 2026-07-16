package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PfsSummary obs并行文件系统输出。
type PfsSummary struct {

	// **参数解释**：obs并行文件系统路径url。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PfsPath string `json:"pfs_path"`
}

func (o PfsSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PfsSummary struct{}"
	}

	return strings.Join([]string{"PfsSummary", string(data)}, " ")
}
