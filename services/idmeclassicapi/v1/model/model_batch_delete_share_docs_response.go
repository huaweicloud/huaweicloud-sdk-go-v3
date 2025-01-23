package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteShareDocsResponse Response Object
type BatchDeleteShareDocsResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。  **默认取值：**  不涉及。
	Result *string `json:"result,omitempty"`

	Data *[]int32 `json:"data,omitempty"`

	// **参数解释：**  异常信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchDeleteShareDocsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteShareDocsResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteShareDocsResponse", string(data)}, " ")
}
