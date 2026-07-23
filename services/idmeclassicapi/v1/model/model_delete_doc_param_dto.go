package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteDocParamDto struct {

	// **参数解释：**  待删除分享权限记录的ID集合。可以通过接口“查询结构化文档分享授权列表”的StructuredDocShareViewDTO.id进行获取。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Ids []string `json:"ids"`
}

func (o DeleteDocParamDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDocParamDto struct{}"
	}

	return strings.Join([]string{"DeleteDocParamDto", string(data)}, " ")
}
