package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateArtifactoryRequest Request Object
type UpdateArtifactoryRequest struct {
	Body *UpdateNotMavenRepoDo `json:"body,omitempty"`
}

func (o UpdateArtifactoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateArtifactoryRequest struct{}"
	}

	return strings.Join([]string{"UpdateArtifactoryRequest", string(data)}, " ")
}
