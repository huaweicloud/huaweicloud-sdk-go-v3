package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagOperationDto struct {

	// **参数解释：**  目标数据实例ID。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	ObjectId string `json:"objectId"`

	// **参数解释：**  当前租户下标签ID。 获取方法请参见[全量数据服务](https://support.huaweicloud.com/usermanual-idme/idme_clientog_0154.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	TagId string `json:"tagId"`

	// **参数解释：**  更新者账号，用于记录执行此操作的用户标识。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Modifier *string `json:"modifier,omitempty"`
}

func (o TagOperationDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagOperationDto struct{}"
	}

	return strings.Join([]string{"TagOperationDto", string(data)}, " ")
}
