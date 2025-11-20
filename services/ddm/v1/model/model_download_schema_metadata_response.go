package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DownloadSchemaMetadataResponse Response Object
type DownloadSchemaMetadataResponse struct {

	// 逻辑库信息。
	CompressedDatabasesInfo *string `json:"compressed_databases_info,omitempty"`

	// 关联的后端DN信息。
	DataNodes      *[]DnInstanceInfo `json:"data_nodes,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o DownloadSchemaMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DownloadSchemaMetadataResponse struct{}"
	}

	return strings.Join([]string{"DownloadSchemaMetadataResponse", string(data)}, " ")
}
