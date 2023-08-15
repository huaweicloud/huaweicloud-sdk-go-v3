package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMetadata2Request Request Object
type ShowMetadata2Request struct {

	// 元数据ID。
	MetadataId string `json:"metadata_id"`
}

func (o ShowMetadata2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMetadata2Request struct{}"
	}

	return strings.Join([]string{"ShowMetadata2Request", string(data)}, " ")
}
