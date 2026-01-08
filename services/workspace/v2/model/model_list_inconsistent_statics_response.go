package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInconsistentStaticsResponse Response Object
type ListInconsistentStaticsResponse struct {

	// 统计productID与桌面池套餐ID不一致的桌面数量
	ProductId *int32 `json:"product_id,omitempty"`

	// 统计imageID与桌面池镜像ID不一致的桌面数量
	ImageId *int32 `json:"image_id,omitempty"`

	// 统计数据盘数量与桌面池配置不一致的桌面数量
	DiskNum *int32 `json:"disk_num,omitempty"`

	// 统计磁盘累加容量与桌面池配置不一致的桌面数量
	DiskSize       *int32 `json:"disk_size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInconsistentStaticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInconsistentStaticsResponse struct{}"
	}

	return strings.Join([]string{"ListInconsistentStaticsResponse", string(data)}, " ")
}
