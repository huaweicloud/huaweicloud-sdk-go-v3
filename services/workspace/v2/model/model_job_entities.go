package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type JobEntities struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 套餐ID。
	ProductId *string `json:"product_id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
