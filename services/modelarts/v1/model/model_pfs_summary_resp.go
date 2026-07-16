package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PfsSummaryResp obs并行文件系统输出。
type PfsSummaryResp struct {

	// **参数解释**：obs并行文件系统路径url。 **取值范围**：不涉及。
	PfsPath string `json:"pfs_path"`
}

func (o PfsSummaryResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PfsSummaryResp struct{}"
	}

	return strings.Join([]string{"PfsSummaryResp", string(data)}, " ")
}
