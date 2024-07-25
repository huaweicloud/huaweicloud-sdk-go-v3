package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueryUsingPostResponse Response Object
type ListQueryUsingPostResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。  **默认取值：**  不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Data *[]PersistableModelQueryViewDto `json:"data,omitempty"`

	// **参数解释：**  异常信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListQueryUsingPostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueryUsingPostResponse struct{}"
	}

	return strings.Join([]string{"ListQueryUsingPostResponse", string(data)}, " ")
}
