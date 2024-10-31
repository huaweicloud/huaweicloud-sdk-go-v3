package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBackupMindmapRequest Request Object
type CreateBackupMindmapRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestBackUpMindmapParam `json:"body,omitempty"`
}

func (o CreateBackupMindmapRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBackupMindmapRequest struct{}"
	}

	return strings.Join([]string{"CreateBackupMindmapRequest", string(data)}, " ")
}
