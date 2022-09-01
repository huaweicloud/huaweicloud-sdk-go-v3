package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListRetirableGrantsRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *ListRetirableGrantsRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ListRetirableGrantsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRetirableGrantsRequest struct{}"
	}

	return strings.Join([]string{"ListRetirableGrantsRequest", string(data)}, " ")
}
