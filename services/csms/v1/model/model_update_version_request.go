package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVersionRequest Request Object
type UpdateVersionRequest struct {

	// 凭据名称。
	SecretName string `json:"secret_name"`

	// 凭据的版本标识符。
	VersionId string `json:"version_id"`

	Body *UpdateVersionRequestBody `json:"body,omitempty"`
}

func (o UpdateVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVersionRequest struct{}"
	}

	return strings.Join([]string{"UpdateVersionRequest", string(data)}, " ")
}
