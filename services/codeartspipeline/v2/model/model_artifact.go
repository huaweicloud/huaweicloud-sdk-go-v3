package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Artifact 产物
type Artifact struct {

	// **参数解释**： 项目ID，获取方式请参见[获取项目ID](https://support.huaweicloud.com/api-codeartsrepo/codehub_api_0014.html)。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 产物名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 产物版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ArtifactVersion *string `json:"artifact_version,omitempty"`

	// **参数解释**： 产物存放平台。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	UploadTarget *string `json:"upload_target,omitempty"`

	// **参数解释**： 产物包类型，例如jar。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ArtifactPackageType *string `json:"artifact_package_type,omitempty"`

	// **参数解释**： 制品仓文件存放路径。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ArtifactUri *string `json:"artifact_uri,omitempty"`

	// **参数解释**： 制品仓下载链接。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ArtifactDownloadUrlWithId *string `json:"artifact_download_url_with_id,omitempty"`

	// **参数解释**： 产物类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ArtifactType *string `json:"artifact_type,omitempty"`

	// **参数解释**： 产物哈希码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	HashCode *[]ArtifactHashCode `json:"hash_code,omitempty"`

	// **参数解释**： 构建任务ID。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 构建任务编号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BuildNo *int32 `json:"build_no,omitempty"`

	// **参数解释**： 构建任务序号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	DailyBuildNumber *string `json:"daily_build_number,omitempty"`

	// **参数解释**： 产物大小。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FileSize *string `json:"file_size,omitempty"`
}

func (o Artifact) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Artifact struct{}"
	}

	return strings.Join([]string{"Artifact", string(data)}, " ")
}
