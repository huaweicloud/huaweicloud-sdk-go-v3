package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateArtifactoryResult struct {
	Artifactory *RepositoryDo `json:"artifactory,omitempty"`
}

func (o CreateArtifactoryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArtifactoryResult struct{}"
	}

	return strings.Join([]string{"CreateArtifactoryResult", string(data)}, " ")
}
