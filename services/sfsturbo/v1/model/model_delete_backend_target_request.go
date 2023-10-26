package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackendTargetRequest Request Object
type DeleteBackendTargetRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统id
	ShareId string `json:"share_id"`

	// 数据存储库 id
	TargetId string `json:"target_id"`

	// 删除后端存储库时是否删除文件系统对应目录文件，默认为 false
	DeleteDataInFileSystem *bool `json:"delete_data_in_file_system,omitempty"`
}

func (o DeleteBackendTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackendTargetRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackendTargetRequest", string(data)}, " ")
}
