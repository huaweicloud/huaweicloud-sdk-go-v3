package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteIpdIssuesRequest Request Object
type BatchDeleteIpdIssuesRequest struct {

	// 项目32位ID，项目唯一标识。通过查询IPD项目列表获取，响应消息体中的id字段的值就是项目ID。
	ProjectId string `json:"project_id"`

	// **参数解释**： 是否永久删除。 **约束限制**： 不涉及。 **取值范围**： - true：彻底删除工作项（适用于回收站中的工作项，彻底删除后不可恢复）。 - false：将工作项移入回收站。 **默认取值**： false。
	IsPermanentDelete *bool `json:"is_permanent_delete,omitempty"`

	// **参数解释**： 当工作项类型为RR或Bug时，工作项的提出项目ID。通过[查询IPD项目列表](ShowIpdProjectList.xml)获取，响应消息体中的**id**字段的值就是项目ID。 **约束限制**： 归属项目和提出项目一致时可不传。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	SrcProjectId *string `json:"src_project_id,omitempty"`

	Body *[]string `json:"body,omitempty"`
}

func (o BatchDeleteIpdIssuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteIpdIssuesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteIpdIssuesRequest", string(data)}, " ")
}
