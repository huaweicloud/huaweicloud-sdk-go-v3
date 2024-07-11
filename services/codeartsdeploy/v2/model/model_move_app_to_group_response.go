package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MoveAppToGroupResponse Response Object
type MoveAppToGroupResponse struct {

	// 移动应用至指定分组结果，仅返回移动失败的列表
	Result *[]MoveAppToGroupResult `json:"result,omitempty"`

	// 请求成功失败状态
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o MoveAppToGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MoveAppToGroupResponse struct{}"
	}

	return strings.Join([]string{"MoveAppToGroupResponse", string(data)}, " ")
}
