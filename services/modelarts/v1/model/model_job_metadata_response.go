package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobMetadataResponse 训练作业元信息。
type JobMetadataResponse struct {

	// **参数解释**：训练作业ID，创建成功后由ModelArts生成返回，无需填写。 **取值范围**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：训练作业名称。 **取值范围**：限制为1-64位只含数字、字母、下划线和中划线的名称。
	Name string `json:"name"`

	// **参数解释**：指定作业所处的工作空间。 **取值范围**：不涉及。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**：对训练作业的描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**：训练作业创建时间戳，单位为毫秒，创建成功后由ModelArts生成返回，无需填写。 **取值范围**：不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**：训练作业创建用户的用户名，创建成功后由ModelArts生成返回，无需填写。 **取值范围**：不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**：训练作业高级功能配置。
	Annotations map[string]string `json:"annotations,omitempty"`

	TrainingExperimentReference *TrainingExperimentResp `json:"training_experiment_reference,omitempty"`
}

func (o JobMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobMetadataResponse struct{}"
	}

	return strings.Join([]string{"JobMetadataResponse", string(data)}, " ")
}
