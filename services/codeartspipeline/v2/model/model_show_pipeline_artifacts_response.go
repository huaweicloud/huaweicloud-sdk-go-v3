package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineArtifactsResponse Response Object
type ShowPipelineArtifactsResponse struct {

	// 产物列表
	Artifacts      *[]Artifact `json:"artifacts,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowPipelineArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineArtifactsResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineArtifactsResponse", string(data)}, " ")
}
