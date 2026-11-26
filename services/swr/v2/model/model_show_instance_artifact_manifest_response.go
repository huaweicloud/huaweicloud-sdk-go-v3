package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceArtifactManifestResponse Response Object
type ShowInstanceArtifactManifestResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceArtifactManifestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceArtifactManifestResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceArtifactManifestResponse", string(data)}, " ")
}
