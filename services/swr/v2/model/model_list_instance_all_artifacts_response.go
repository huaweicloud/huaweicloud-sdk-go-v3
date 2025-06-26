package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceAllArtifactsResponse Response Object
type ListInstanceAllArtifactsResponse struct {

	// 制品列表
	Artifacts *[]Artifact `json:"artifacts,omitempty"`

	// 下一次分页查询的起始ID。如果未返回该值，则表示数据已查询完毕
	NextMarker     *int32 `json:"next_marker,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceAllArtifactsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceAllArtifactsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceAllArtifactsResponse", string(data)}, " ")
}
