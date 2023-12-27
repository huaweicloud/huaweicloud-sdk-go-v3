package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchArtifactsRequest Request Object
type SearchArtifactsRequest struct {
	Body *IdeRepoSearchDo `json:"body,omitempty"`
}

func (o SearchArtifactsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchArtifactsRequest struct{}"
	}

	return strings.Join([]string{"SearchArtifactsRequest", string(data)}, " ")
}
