package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProductResponse Response Object
type ListProductResponse struct {

	// 可用分区。将服务创建到指定的可用分区。如果不指定则使用系统随机的可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 系统类型，当前仅支持Windows。 * `Linux` - * `Windows` - * `Other` -
	OsType *string `json:"os_type,omitempty"`

	// 产品列表
	Products       *[]ProductInfo `json:"products,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListProductResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProductResponse struct{}"
	}

	return strings.Join([]string{"ListProductResponse", string(data)}, " ")
}
