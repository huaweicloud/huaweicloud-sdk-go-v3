package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserPrivilegesRequest Request Object
type ListUserPrivilegesRequest struct {

	// **参数解释**： 项目ID，可以从调用API处获取，也可以从控制台获取。获取方式请参考[获取项目ID](CloudArtifact_api_0015.xml)。 **约束限制**： 不涉及。 **取值范围**： 只能由英文字母、数字组成，且长度为32个字符。 **默认取值**： 不涉及。
	ProjectId string `json:"project_id"`
}

func (o ListUserPrivilegesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserPrivilegesRequest struct{}"
	}

	return strings.Join([]string{"ListUserPrivilegesRequest", string(data)}, " ")
}
