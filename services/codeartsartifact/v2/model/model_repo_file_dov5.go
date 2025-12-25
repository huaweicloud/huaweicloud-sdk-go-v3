package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RepoFileDov5 struct {

	// **参数解释**: id。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 文件id。 **取值范围**: 不涉及。
	FileId *string `json:"file_id,omitempty"`

	// **参数解释**: 仓库id。 **取值范围**: 不涉及。
	RepoName *string `json:"repo_name,omitempty"`

	// **参数解释**: 项目名称。 **取值范围**: 不涉及。
	ProjectName *string `json:"project_name,omitempty"`

	// **参数解释**: 父级目录ID。 **取值范围**: 不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**: 文件名。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 文件名。 **取值范围**: 不涉及。
	FileName *string `json:"file_name,omitempty"`

	// **参数解释**: 文件类型, folder代表是目录,file代表文件。 **取值范围**: 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**: 发布包状态 test为测试包 prod为发布包。 **取值范围**: 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**: 文件扩展名。 **取值范围**: 不涉及。
	Extension *string `json:"extension,omitempty"`

	// **参数解释**: 文件路径。 **取值范围**: 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释**: 文件路径(含项目)。 **取值范围**: 不涉及。
	FullPath *string `json:"full_path,omitempty"`

	// **参数解释**: 文件大小,单位为byte。 **取值范围**: 不涉及。
	Size *int64 `json:"size,omitempty"`

	// **参数解释**: md5校验和。 **取值范围**: 不涉及。
	Md5 *string `json:"md5,omitempty"`

	// **参数解释**: sha256校验和。 **取值范围**: 不涉及。
	Sha256 *string `json:"sha256,omitempty"`

	// **参数解释**: 下载地址。 **取值范围**: 不涉及。
	DownloadUrl *string `json:"download_url,omitempty"`

	// **参数解释**: 带有id的下载地址。 **取值范围**: 不涉及。
	DownloadUrlWithId *string `json:"download_url_with_id,omitempty"`

	// **参数解释**: 历史版本遗留字段，请忽略。 **取值范围**: 不涉及。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释**: 判断当前文件或文件夹父目录是否为版本路径，即仓库下第一层子目录。 **取值范围**: - true：父目录是版本路径。 - false：父目录不是版本路径。
	VersionEnable *bool `json:"version_enable,omitempty"`

	// **参数解释**: migrated_state。 **取值范围**: 该参数为内部数据改造参数，无业务含义，请忽略。
	MigratedState *int32 `json:"migrated_state,omitempty"`

	// **参数解释**: 该参数无返回值，请忽略。 **取值范围**: 不涉及。
	UploadId *string `json:"upload_id,omitempty"`
}

func (o RepoFileDov5) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RepoFileDov5 struct{}"
	}

	return strings.Join([]string{"RepoFileDov5", string(data)}, " ")
}
