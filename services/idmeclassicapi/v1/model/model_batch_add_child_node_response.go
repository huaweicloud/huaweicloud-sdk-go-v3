package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddChildNodeResponse Response Object
type BatchAddChildNodeResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据，返回批量添加后的子节点信息列表，包含更新后的根节点、父节点、全路径等树形结构属性。  **取值范围：**  不涉及。
	Data *[]TreeableModelViewDto `json:"data,omitempty"`

	// **参数解释：**  异常信息，当请求失败时返回具体的错误描述。  **取值范围：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchAddChildNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddChildNodeResponse struct{}"
	}

	return strings.Join([]string{"BatchAddChildNodeResponse", string(data)}, " ")
}
