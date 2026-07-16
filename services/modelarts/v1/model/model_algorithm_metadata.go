package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlgorithmMetadata 算法的元数据，描述算法基本信息。
type AlgorithmMetadata struct {

	// 算法uuid，创建算法时无需填写。
	Id *int32 `json:"id,omitempty"`

	// 算法名称。限制为1-64位只含数字、字母、下划线和中划线的名称。
	Name string `json:"name"`

	// 对算法的描述，默认为“NULL”，字符串的长度限制为[0, 256]。
	Description *string `json:"description,omitempty"`

	// 指定算法所处的工作空间，默认值为“0”。“0” 为默认的工作空间。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 指定算法所属的ai项目，默认值为\"default-ai-project\"。ai项目已下线，无需关注。
	AiProject *string `json:"ai_project,omitempty"`
}

func (o AlgorithmMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlgorithmMetadata struct{}"
	}

	return strings.Join([]string{"AlgorithmMetadata", string(data)}, " ")
}
