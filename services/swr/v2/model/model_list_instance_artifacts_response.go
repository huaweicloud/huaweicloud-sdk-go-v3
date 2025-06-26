package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceArtifactsResponse Response Object
type ListInstanceArtifactsResponse struct {

	// 制品列表
	Artifacts *[]Artifact `json:"artifacts,omitempty"`

	// 制品总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceArtifactsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceArtifactsResponse", string(data)}, " ")
}
