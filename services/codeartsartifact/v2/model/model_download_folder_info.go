package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DownloadFolderInfo struct {

	// **参数解释**： 文件总数。  **取值范围**： 不涉及。
	FilesCount *int64 `json:"filesCount,omitempty"`

	// **参数解释**： 占用空间。  **取值范围**： 不涉及。
	UsedSpace *string `json:"usedSpace,omitempty"`
}

func (o DownloadFolderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadFolderInfo struct{}"
	}

	return strings.Join([]string{"DownloadFolderInfo", string(data)}, " ")
}
