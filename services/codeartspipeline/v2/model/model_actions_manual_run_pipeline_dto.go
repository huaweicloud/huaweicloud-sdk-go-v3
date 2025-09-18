package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionsManualRunPipelineDto struct {

	// **参数解释**： 触发URL。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	HttpsUrl *string `json:"https_url,omitempty"`

	// **参数解释**： 文件存储路径。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释**： 文件详情。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	FileContent *string `json:"file_content,omitempty"`

	// **参数解释**： 代码源分支名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Branch *string `json:"branch,omitempty"`

	// **参数解释**： 文件编码方式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Encoding *string `json:"encoding,omitempty"`

	// **参数解释**： 代码源标签。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释**： 代码源提交记录id。 **约束限制**： 不涉及。 **取值范围**： 40位字符。 **默认取值**： 不涉及。
	CommitId *string `json:"commit_id,omitempty"`

	// **参数解释**： 代码源访问权限token。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	AccessToken *string `json:"access_token,omitempty"`
}

func (o ActionsManualRunPipelineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionsManualRunPipelineDto struct{}"
	}

	return strings.Join([]string{"ActionsManualRunPipelineDto", string(data)}, " ")
}
