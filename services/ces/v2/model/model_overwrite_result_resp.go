package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OverwriteResultResp **参数解释**： 如果出现同名模板，返回是否覆盖的结果信息。
type OverwriteResultResp struct {

	// **参数解释**：已覆盖的模板ID值。 **取值范围**：长度为[2,64]个字符。
	SuccessIds *string `json:"success_ids,omitempty"`

	// **参数解释**：未覆盖的模板ID值。 **取值范围**：长度为[2,64]个字符。
	FailedIds *string `json:"failed_ids,omitempty"`
}

func (o OverwriteResultResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OverwriteResultResp struct{}"
	}

	return strings.Join([]string{"OverwriteResultResp", string(data)}, " ")
}
