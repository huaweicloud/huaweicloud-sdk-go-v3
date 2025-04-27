package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublicZoneLinesRequest Request Object
type ListPublicZoneLinesRequest struct {

	// 域名ID。
	ZoneId string `json:"zone_id"`

	// 每页返回的资源个数。  取值范围：0~500  取值一般为10，20，50。默认值为500。
	Limit *int32 `json:"limit,omitempty"`

	// 分页查询起始偏移量，表示从偏移量的下一个资源开始查询。  取值范围：0~2147483647  默认值为0。  当前设置marker不为空时，以marker为分页起始标识。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPublicZoneLinesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicZoneLinesRequest struct{}"
	}

	return strings.Join([]string{"ListPublicZoneLinesRequest", string(data)}, " ")
}
