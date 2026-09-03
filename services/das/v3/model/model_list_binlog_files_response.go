package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBinlogFilesResponse Response Object
type ListBinlogFilesResponse struct {

	// 总数
	Total *int32 `json:"total,omitempty"`

	// binlog文件列表
	FileList       *[]BinlogFileInfo `json:"file_list,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListBinlogFilesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBinlogFilesResponse struct{}"
	}

	return strings.Join([]string{"ListBinlogFilesResponse", string(data)}, " ")
}
