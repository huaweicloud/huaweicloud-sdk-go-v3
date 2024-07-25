package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Result **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。  **默认取值：**  不涉及。
type Result struct {
}

func (o Result) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Result struct{}"
	}

	return strings.Join([]string{"Result", string(data)}, " ")
}
