package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueAccessoryV2 struct {

	// **参数解释：** 附件id。 **取值范围：** 不涉及。
	AttachmentId *int32 `json:"attachment_id,omitempty"`

	// **参数解释：** 工作项数字id。 **取值范围：** 不涉及。
	IssueId *int32 `json:"issue_id,omitempty"`

	// **参数解释：** 附件的上传者数字ID。 **取值范围：** 不涉及。
	CreatorNumId *int32 `json:"creator_num_id,omitempty"`

	// **参数解释：** 附件创建时间（示例：2025-08-04 19:43:46）。 **取值范围：** 不涉及。
	CreatedDate *string `json:"created_date,omitempty"`

	// **参数解释：** 附件名称。 **取值范围：** 不涉及。
	FileName *string `json:"file_name,omitempty"`

	// **参数解释：** 附件所属类型。 **取值范围：** scrum。
	ContainerType *string `json:"container_type,omitempty"`

	// **参数解释：** 附件在服务器上实际名称。 **取值范围：** 不涉及。
	DiskFileName *string `json:"disk_file_name,omitempty"`

	// **参数解释：** 附件来源。 **取值范围：** 1：工作项本地上传的文档；2：关联的文档。
	Digest *string `json:"digest,omitempty"`

	// **参数解释：** 附件在服务器上的路径。 **取值范围：** 不涉及。
	DiskDirectory *string `json:"disk_directory,omitempty"`

	// **参数解释：** 附件的上传者uuid。 **取值范围：** 不涉及。
	CreatorId *string `json:"creator_id,omitempty"`
}

func (o IssueAccessoryV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueAccessoryV2 struct{}"
	}

	return strings.Join([]string{"IssueAccessoryV2", string(data)}, " ")
}
