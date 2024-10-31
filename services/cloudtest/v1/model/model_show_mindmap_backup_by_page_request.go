package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapBackupByPageRequest Request Object
type ShowMindmapBackupByPageRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	Body *CommRequestMindmapBackupPageParam `json:"body,omitempty"`
}

func (o ShowMindmapBackupByPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapBackupByPageRequest struct{}"
	}

	return strings.Join([]string{"ShowMindmapBackupByPageRequest", string(data)}, " ")
}
