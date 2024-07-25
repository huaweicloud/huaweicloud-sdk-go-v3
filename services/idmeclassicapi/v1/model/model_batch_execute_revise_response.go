package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchExecuteReviseResponse Response Object
type BatchExecuteReviseResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。  **默认取值：**  不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Data *[]VersionModelViewDto `json:"data,omitempty"`

	// **参数解释：**  异常信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o BatchExecuteReviseResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchExecuteReviseResponse struct{}"
	}

	return strings.Join([]string{"BatchExecuteReviseResponse", string(data)}, " ")
}
