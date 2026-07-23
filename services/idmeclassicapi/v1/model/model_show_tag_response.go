package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagResponse Response Object
type ShowTagResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据，返回数据实例绑定的标签详情列表。若实例未绑定标签，则返回空数组。  **取值范围：**  不涉及。
	Data *[]interface{} `json:"data,omitempty"`

	// **参数解释：**  异常信息，当请求失败时返回具体的错误描述。  **取值范围：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagResponse struct{}"
	}

	return strings.Join([]string{"ShowTagResponse", string(data)}, " ")
}
