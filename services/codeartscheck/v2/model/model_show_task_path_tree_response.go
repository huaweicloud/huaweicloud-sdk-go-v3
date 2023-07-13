package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTaskPathTreeResponse Response Object
type ShowTaskPathTreeResponse struct {

	// 任务的目录树信息
	Info *[]TreeNode `json:"info,omitempty"`

	// 数目
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowTaskPathTreeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTaskPathTreeResponse struct{}"
	}

	return strings.Join([]string{"ShowTaskPathTreeResponse", string(data)}, " ")
}
