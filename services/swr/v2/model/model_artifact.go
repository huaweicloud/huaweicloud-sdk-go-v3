package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Artifact struct {

	// 制品摘要
	Digest *string `json:"digest,omitempty"`

	// 制品ID
	Id *int32 `json:"id,omitempty"`

	// 仓库ID
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 仓库名称
	RepositoryName *string `json:"repository_name,omitempty"`

	// 制品类型
	Type *string `json:"type,omitempty"`

	// 命名空间ID
	NamespaceId *int32 `json:"namespace_id,omitempty"`

	// 制品MIME类型
	MediaType *string `json:"media_type,omitempty"`

	// 制品元数据MIME类型
	ManifestMediaType *string `json:"manifest_media_type,omitempty"`

	// 最近一次拉取时间
	PullTime *string `json:"pull_time,omitempty"`

	// 最近一次上传时间
	PushTime *string `json:"push_time,omitempty"`

	// 制品大小，单位：Byte
	Size *int32 `json:"size,omitempty"`

	// 制品版本的Tag列表
	Tags *[]ArtifactTag `json:"tags,omitempty"`
}

func (o Artifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Artifact struct{}"
	}

	return strings.Join([]string{"Artifact", string(data)}, " ")
}
