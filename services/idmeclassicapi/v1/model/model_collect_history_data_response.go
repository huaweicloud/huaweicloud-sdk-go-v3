package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectHistoryDataResponse Response Object
type CollectHistoryDataResponse struct {

	// **参数解释：**  请求结果。  **取值范围：**  - SUCCESS：请求成功。 - FAIL：请求失败。  **默认取值：**  不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释：**  请求数据。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Data *[]StatisticsRvo `json:"data,omitempty"`

	// **参数解释：**  异常信息。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Errors         *[]string `json:"errors,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o CollectHistoryDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectHistoryDataResponse struct{}"
	}

	return strings.Join([]string{"CollectHistoryDataResponse", string(data)}, " ")
}
