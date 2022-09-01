package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateKeyAliasRequest struct {

	// API版本号
	VersionId string `json:"version_id" xml:"version_id"`

	Body *UpdateKeyAliasRequestBody `json:"body,omitempty" xml:"body"`
}

func (o UpdateKeyAliasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKeyAliasRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyAliasRequest", string(data)}, " ")
}
