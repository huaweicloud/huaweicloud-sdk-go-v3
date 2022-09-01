package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CancelSelfGrantRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *RevokeGrantRequestBody `json:"body,omitempty" xml:"body"`
}

func (o CancelSelfGrantRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelSelfGrantRequest struct{}"
	}

	return strings.Join([]string{"CancelSelfGrantRequest", string(data)}, " ")
}
