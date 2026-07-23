package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryShareDocsResponse Response Object
type ListQueryShareDocsResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据，返回目标结构化文档的分享授权记录列表。 每条记录包含分享人、被分享人、权限类型等详细信息。  **取值范围：**  不涉及。
	Data *[]StructuredDocShareViewDto `json:"data,omitempty"`

	// **参数解释：**  异常信息，当请求失败时返回具体的错误描述。  **取值范围：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListQueryShareDocsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryShareDocsResponse struct{}"
	}

	return strings.Join([]string{"ListQueryShareDocsResponse", string(data)}, " ")
}
