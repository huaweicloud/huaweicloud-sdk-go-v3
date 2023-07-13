package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGraphMetadatasRequest Request Object
type ListGraphMetadatasRequest struct {

	// 元数据ID。
	MetadataId string `json:"metadata_id"`
}

func (o ListGraphMetadatasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGraphMetadatasRequest struct{}"
	}

	return strings.Join([]string{"ListGraphMetadatasRequest", string(data)}, " ")
}
