package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepositoryTag struct {

	// 版本名称
	Name *string `json:"name,omitempty"`

	// 关联的制品ID
	ArtifactId *int32 `json:"artifact_id,omitempty"`

	// 关联的制品摘要
	Digest *string `json:"digest,omitempty"`

	// 版本ID
	Id *int32 `json:"id,omitempty"`

	// 制品仓库仓库ID
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 制品类型，比如IMAGE
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
	Size *int64 `json:"size,omitempty"`
}

func (o RepositoryTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepositoryTag struct{}"
	}

	return strings.Join([]string{"RepositoryTag", string(data)}, " ")
}
