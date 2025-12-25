package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteNodeResponse Response Object
type DeleteNodeResponse struct {

	// 删除失败列表
	DeleteFailList *[]Node `json:"delete_fail_list,omitempty"`

	// 删除成功列表
	DeleteSuccessList *[]Node `json:"delete_success_list,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o DeleteNodeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteNodeResponse struct{}"
	}

	return strings.Join([]string{"DeleteNodeResponse", string(data)}, " ")
}
