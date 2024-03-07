package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Artifact 产物
type Artifact struct {

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 版本
	ArtifactVersion *string `json:"artifact_version,omitempty"`

	// 存放平台类型
	UploadTarget *string `json:"upload_target,omitempty"`

	// 产物包类型
	ArtifactPackageType *string `json:"artifact_package_type,omitempty"`

	// 制品仓路径
	ArtifactUri *string `json:"artifact_uri,omitempty"`

	// 制品仓下载链接
	ArtifactDownloadUrlWithId *string `json:"artifact_download_url_with_id,omitempty"`

	// 产物类型
	ArtifactType *string `json:"artifact_type,omitempty"`

	// 哈希码
	HashCode *[]ArtifactHashCode `json:"hash_code,omitempty"`

	// 构建任务ID
	JobId *string `json:"job_id,omitempty"`

	// 构建任务编号
	BuildNo *int32 `json:"build_no,omitempty"`

	// 构建任务序号
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`

	// 产物大小
	FileSize *string `json:"file_size,omitempty"`
}

func (o Artifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Artifact struct{}"
	}

	return strings.Join([]string{"Artifact", string(data)}, " ")
}
