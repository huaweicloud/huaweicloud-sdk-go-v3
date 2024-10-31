package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMindmapBackupByIdRequest Request Object
type ShowMindmapBackupByIdRequest struct {

	// 项目ID，固定长度32位字符（字母和数字）。
	ProjectId string `json:"project_id"`

	// 备份ID
	Id string `json:"id"`
}

func (o ShowMindmapBackupByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMindmapBackupByIdRequest struct{}"
	}

	return strings.Join([]string{"ShowMindmapBackupByIdRequest", string(data)}, " ")
}
