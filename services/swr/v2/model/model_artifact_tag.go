package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ArtifactTag struct {

	// Tag ID
	Id *int64 `json:"id,omitempty"`

	// 制品仓库ID
	RepositoryId *int64 `json:"repository_id,omitempty"`

	// 制品版本ID
	ArtifactId *int64 `json:"artifact_id,omitempty"`

	// tag名称
	Name *string `json:"name,omitempty"`

	// tag的上传时间
	PushTime *sdktime.SdkTime `json:"push_time,omitempty"`

	// tag的下载时间
	PullTime *sdktime.SdkTime `json:"pull_time,omitempty"`
}

func (o ArtifactTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ArtifactTag struct{}"
	}

	return strings.Join([]string{"ArtifactTag", string(data)}, " ")
}
