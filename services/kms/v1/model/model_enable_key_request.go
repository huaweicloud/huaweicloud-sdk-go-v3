package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type EnableKeyRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *OperateKeyRequestBody `json:"body,omitempty" xml:"body"`
}

func (o EnableKeyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableKeyRequest struct{}"
	}

	return strings.Join([]string{"EnableKeyRequest", string(data)}, " ")
}
