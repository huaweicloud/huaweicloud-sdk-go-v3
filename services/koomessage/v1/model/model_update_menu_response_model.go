package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateMenuResponseModel struct {

	// 菜单ID。
	MenuId *string `json:"menu_id,omitempty"`

	// 操作记录ID。
	LogId *string `json:"log_id,omitempty"`
}

func (o UpdateMenuResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMenuResponseModel struct{}"
	}

	return strings.Join([]string{"UpdateMenuResponseModel", string(data)}, " ")
}
