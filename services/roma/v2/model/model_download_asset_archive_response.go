package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DownloadAssetArchiveResponse struct {
	// 应用列表

	Apps *[]string `json:"apps,omitempty"`
	// 任务列表

	Tasks          *[]AppAssetTasks `json:"tasks,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o DownloadAssetArchiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetArchiveResponse struct{}"
	}

	return strings.Join([]string{"DownloadAssetArchiveResponse", string(data)}, " ")
}
