package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteApplicationGroupRequest Request Object
type DeleteApplicationGroupRequest struct {

	// 分组id。
	Id string `json:"id"`

	// 是否强制删除。
	ForceDelete *bool `json:"force_delete,omitempty"`
}

func (o DeleteApplicationGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteApplicationGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteApplicationGroupRequest", string(data)}, " ")
}
