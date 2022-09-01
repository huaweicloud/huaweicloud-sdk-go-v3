package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListKeysRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *ListKeysRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ListKeysRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKeysRequest struct{}"
	}

	return strings.Join([]string{"ListKeysRequest", string(data)}, " ")
}
