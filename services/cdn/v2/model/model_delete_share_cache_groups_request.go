package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteShareCacheGroupsRequest Request Object
type DeleteShareCacheGroupsRequest struct {

	// **参数解释：** 共享缓存组ID，可通过查询共享缓存组列表接口获取 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Id string `json:"id"`
}

func (o DeleteShareCacheGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteShareCacheGroupsRequest struct{}"
	}

	return strings.Join([]string{"DeleteShareCacheGroupsRequest", string(data)}, " ")
}
