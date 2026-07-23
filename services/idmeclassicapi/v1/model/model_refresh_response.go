package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RefreshResponse Response Object
type RefreshResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据，返回刷新任务的相关信息。 由于本接口为异步接口，返回的data不代表刷新已完成，仅表示任务已提交。  **取值范围：**  不涉及。
	Data *[]interface{} `json:"data,omitempty"`

	// **参数解释：**  异常信息，当请求失败时返回具体的错误描述。  **取值范围：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o RefreshResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RefreshResponse struct{}"
	}

	return strings.Join([]string{"RefreshResponse", string(data)}, " ")
}
