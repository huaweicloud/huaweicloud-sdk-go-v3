package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PushMenuInfoRequest Request Object
type PushMenuInfoRequest struct {

	// 菜单ID。
	MenuId string `json:"menu_id"`
}

func (o PushMenuInfoRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PushMenuInfoRequest struct{}"
	}

	return strings.Join([]string{"PushMenuInfoRequest", string(data)}, " ")
}
