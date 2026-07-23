package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdsDto struct {

	// **参数解释：**  数据实例ID列表。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Ids []string `json:"ids"`
}

func (o PersistObjectIdsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdsDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdsDto", string(data)}, " ")
}
