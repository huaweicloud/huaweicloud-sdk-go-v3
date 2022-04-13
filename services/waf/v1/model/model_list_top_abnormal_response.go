package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListTopAbnormalResponse struct {
	// 攻击类型种类

	Total *int32 `json:"total,omitempty"`
	// CountItem详细信息

	Items          *[]UrlCountItem `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTopAbnormalResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopAbnormalResponse struct{}"
	}

	return strings.Join([]string{"ListTopAbnormalResponse", string(data)}, " ")
}
