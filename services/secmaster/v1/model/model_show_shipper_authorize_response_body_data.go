package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowShipperAuthorizeResponseBodyData 结果数据
type ShowShipperAuthorizeResponseBodyData struct {

	// 结果数据
	Data *[]ShowShipperAuthorizeResponseBodyDataData `json:"data,omitempty"`

	// 每页显示的记录数
	Limit *int32 `json:"limit,omitempty"`

	// 当前页的起始偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 总记录数
	Total *int32 `json:"total,omitempty"`
}

func (o ShowShipperAuthorizeResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShipperAuthorizeResponseBodyData struct{}"
	}

	return strings.Join([]string{"ShowShipperAuthorizeResponseBodyData", string(data)}, " ")
}
