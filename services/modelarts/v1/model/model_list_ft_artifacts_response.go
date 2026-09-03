package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFtArtifactsResponse Response Object
type ListFtArtifactsResponse struct {

	// 本次查询到的数据条目数。
	Total *int32 `json:"total,omitempty"`

	// 产物信息。
	ArtifactInfo   *[]ArtifactInfo `json:"artifact_info,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListFtArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFtArtifactsResponse struct{}"
	}

	return strings.Join([]string{"ListFtArtifactsResponse", string(data)}, " ")
}
