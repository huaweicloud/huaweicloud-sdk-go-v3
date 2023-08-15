package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSecretVersionResponse Response Object
type CreateSecretVersionResponse struct {
	VersionMetadata *VersionMetadata `json:"version_metadata,omitempty"`
	HttpStatusCode  int              `json:"-"`
}

func (o CreateSecretVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecretVersionResponse struct{}"
	}

	return strings.Join([]string{"CreateSecretVersionResponse", string(data)}, " ")
}
