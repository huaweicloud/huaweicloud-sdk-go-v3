package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteMetadata2Request Request Object
type DeleteMetadata2Request struct {

	// 元数据ID。
	MetadataId string `json:"metadata_id"`
}

func (o DeleteMetadata2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMetadata2Request struct{}"
	}

	return strings.Join([]string{"DeleteMetadata2Request", string(data)}, " ")
}
