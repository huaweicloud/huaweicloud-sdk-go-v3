package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolDesktopsDetailRequest Request Object
type ListPoolDesktopsDetailRequest struct {

	// 桌面池ID。
	PoolId string `json:"pool_id"`

	// 通过该类型过滤出与桌面池规格类型不一致的桌面 - PRODUCT: 查找productID与桌面池套餐ID不一致的桌面 - IMAGE: 查找imageID与桌面池镜像ID不一致的桌面
	InconsistentType *string `json:"inconsistent_type,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，返回桌面数量限制。取值范围0-100，默认值是10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListPoolDesktopsDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolDesktopsDetailRequest struct{}"
	}

	return strings.Join([]string{"ListPoolDesktopsDetailRequest", string(data)}, " ")
}
