package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteMetadataRequest struct {
	// 元数据ID。

	MetadataId string `json:"metadata_id"`
}

func (o DeleteMetadataRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMetadataRequest struct{}"
	}

	return strings.Join([]string{"DeleteMetadataRequest", string(data)}, " ")
}
