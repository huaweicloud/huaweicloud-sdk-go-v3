package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CompareVersionVo struct {

	// **参数解释：**  数据实例ID，用于指定待对比版本的数据实例。 获取方法请参见[分页查询实例 - ShowFindUsingPost](https://support.huaweicloud.com/api-idme/ShowFindUsingPost.html)。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`

	// **参数解释：**  基础版本号，作为对比的基准版本。对应系统版本号（rdmVersion）。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	BasicVersion string `json:"basicVersion"`

	// **参数解释：**  对比版本号，与基础版本进行对比的目标版本。对应系统版本号（rdmVersion）。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CorrelationVersion string `json:"correlationVersion"`
}

func (o CompareVersionVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompareVersionVo struct{}"
	}

	return strings.Join([]string{"CompareVersionVo", string(data)}, " ")
}
