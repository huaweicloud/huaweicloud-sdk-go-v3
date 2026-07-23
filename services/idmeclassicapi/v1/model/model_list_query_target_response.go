package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryTargetResponse Response Object
type ListQueryTargetResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  查询到的目标模型数据实例列表。返回的实例信息包含目标模型的“列表属性”（即模型中标记为列表展示的属性）。  - 若目标模型存在“参考对象”属性且参考抽象模型，该属性仅返回模型英文名称和ID。 - 若参考实体模型，该属性返回空值。  **取值范围：**  不涉及。
	Data *[]StudentQueryViewDto `json:"data,omitempty"`

	// **参数解释：**  异常信息，当请求失败时返回具体的错误描述。  **取值范围：**  不涉及。
	Errors *[]string `json:"errors,omitempty"`

	PageInfo       *PageInfoViewDto `json:"pageInfo,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListQueryTargetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryTargetResponse struct{}"
	}

	return strings.Join([]string{"ListQueryTargetResponse", string(data)}, " ")
}
