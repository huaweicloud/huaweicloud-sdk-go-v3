package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PersistObjectIdDto struct {

	// **参数解释：**  数据实例ID，用于指定待查询的数据实例。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`
}

func (o PersistObjectIdDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PersistObjectIdDto struct{}"
	}

	return strings.Join([]string{"PersistObjectIdDto", string(data)}, " ")
}
