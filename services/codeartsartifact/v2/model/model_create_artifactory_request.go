package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateArtifactoryRequest Request Object
type CreateArtifactoryRequest struct {
	Body *CreateNotMavenRepoDo `json:"body,omitempty"`
}

func (o CreateArtifactoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArtifactoryRequest struct{}"
	}

	return strings.Join([]string{"CreateArtifactoryRequest", string(data)}, " ")
}
