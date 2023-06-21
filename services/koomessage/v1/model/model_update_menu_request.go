package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateMenuRequest struct {

	// 菜单ID。
	MenuId string `json:"menu_id"`

	Body *UpdateMenuRequestBody `json:"body,omitempty"`
}

func (o UpdateMenuRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMenuRequest struct{}"
	}

	return strings.Join([]string{"UpdateMenuRequest", string(data)}, " ")
}
