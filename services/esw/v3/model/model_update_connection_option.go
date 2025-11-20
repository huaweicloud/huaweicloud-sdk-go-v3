package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConnectionOption struct {

	// - 参数解释：二层连接的名称。 - 约束限制：   - 长度范围为1~64个字符。   - 名称由中文、英文字母、数字、下划线（_）、中划线（-）、点（.）组成。 - 取值范围：不涉及。 - 默认取值：不涉及。
	Name string `json:"name"`
}

func (o UpdateConnectionOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConnectionOption struct{}"
	}

	return strings.Join([]string{"UpdateConnectionOption", string(data)}, " ")
}
