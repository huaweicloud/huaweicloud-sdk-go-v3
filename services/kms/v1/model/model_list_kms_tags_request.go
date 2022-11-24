package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListKmsTagsRequest struct {

	// API版本号
	VersionId string `json:"version_id"`
}

func (o ListKmsTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsTagsRequest struct{}"
	}

	return strings.Join([]string{"ListKmsTagsRequest", string(data)}, " ")
}
