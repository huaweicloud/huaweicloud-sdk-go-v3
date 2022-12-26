package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDashBoardReqBody struct {

	// 仪表盘分组名称
	GroupName *string `json:"group_name,omitempty"`

	// 仪表盘名称
	Title string `json:"title"`
}

func (o CreateDashBoardReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDashBoardReqBody struct{}"
	}

	return strings.Join([]string{"CreateDashBoardReqBody", string(data)}, " ")
}
