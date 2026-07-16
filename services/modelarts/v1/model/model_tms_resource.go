package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TmsResource 根据条件查询得到的TMS返回数据结构。
type TmsResource struct {

	// **参数解释：** 资源详情，用于扩展，默认为空。 **取值范围：** 不涉及
	ResourceDetail *interface{} `json:"resource_detail"`

	// **参数解释：** 资源ID。 **取值范围：** 不涉及
	ResourceId string `json:"resource_id"`

	// **参数解释：** 资源名称。 **取值范围：** 不涉及
	ResourceName string `json:"resource_name"`

	// **参数解释：** 当前资源的所有标签。 **取值范围：** 不涉及
	Tags []InferTmsTag `json:"tags"`
}

func (o TmsResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TmsResource struct{}"
	}

	return strings.Join([]string{"TmsResource", string(data)}, " ")
}
