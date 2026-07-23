package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryParentDto struct {

	// **参数解释：**  子节点实例ID，用于指定待查询所有父节点的目标数据实例。 例如BOM中的某个子装配节点ID、组织架构中的某个部门ID等。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	ChildId string `json:"childId"`
}

func (o QueryParentDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryParentDto struct{}"
	}

	return strings.Join([]string{"QueryParentDto", string(data)}, " ")
}
