package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishFtArtifactsResponse Response Object
type PublishFtArtifactsResponse struct {

	// 模型ID
	ModelId        *string `json:"model_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o PublishFtArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishFtArtifactsResponse struct{}"
	}

	return strings.Join([]string{"PublishFtArtifactsResponse", string(data)}, " ")
}
