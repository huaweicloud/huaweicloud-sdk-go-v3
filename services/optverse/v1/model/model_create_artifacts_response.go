package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateArtifactsResponse Response Object
type CreateArtifactsResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateArtifactsResponse struct{}"
	}

	return strings.Join([]string{"CreateArtifactsResponse", string(data)}, " ")
}
