package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DownloadAssetArchiveRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// 作业执行成功后，查询作业进度返回的entities.archive_id字段
	ArchiveId string `json:"archive_id" xml:"archive_id"`
}

func (o DownloadAssetArchiveRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadAssetArchiveRequest struct{}"
	}

	return strings.Join([]string{"DownloadAssetArchiveRequest", string(data)}, " ")
}
