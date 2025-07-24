package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteBackendTargetRequest Request Object
type DeleteBackendTargetRequest struct {

	// MIME类型
	ContentType string `json:"Content-Type"`

	// 文件系统ID
	ShareId string `json:"share_id"`

	// 绑定关系ID
	TargetId string `json:"target_id"`

	// 删除后端存储时是否同时删除文件系统内的联动目录及其数据文件，默认为 false。数据删除后无法恢复，请谨慎操作。
	DeleteDataInFileSystem *bool `json:"delete_data_in_file_system,omitempty"`
}

func (o DeleteBackendTargetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteBackendTargetRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackendTargetRequest", string(data)}, " ")
}
