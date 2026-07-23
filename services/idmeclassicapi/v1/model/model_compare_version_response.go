package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompareVersionResponse Response Object
type CompareVersionResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  版本对比数据。data[0]为基础版本对象，data[1]为当前版本对象和基础版本对象的差异数据。  **取值范围：**  不涉及。
	Data *[]interface{} `json:"data,omitempty"`

	// **参数解释：**  异常信息，当请求失败时返回具体的错误描述。  **取值范围：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CompareVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionResponse struct{}"
	}

	return strings.Join([]string{"CompareVersionResponse", string(data)}, " ")
}
