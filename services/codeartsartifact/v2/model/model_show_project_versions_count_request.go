package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectVersionsCountRequest Request Object
type ShowProjectVersionsCountRequest struct {

	// **参数解释**： 项目ID，可以调用API获取，也可以从控制台获取，获取方式请参考[项目ID](CloudArtifact_api_0015.xml)。 **约束限制**： 不涉及。 **取值范围**： 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`

	// **参数解释**： 发布库版本的名称。 **约束限制**： 不涉及。 **取值范围**： 英文字母、数字、特殊字符支持中划线、下划线和英文句号，长度为1-128个字符。 **默认取值**： 不涉及。
	BuildVersion *string `json:"build_version,omitempty"`
}

func (o ShowProjectVersionsCountRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectVersionsCountRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectVersionsCountRequest", string(data)}, " ")
}
