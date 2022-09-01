package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 作业对应服务的标题
type TaskDetailsServiceTitle struct {

	// 作业对应服务的中文标题
	Zh *string `json:"zh,omitempty" xml:"zh"`

	// 作业对应服务的英文标题
	En *string `json:"en,omitempty" xml:"en"`
}

func (o TaskDetailsServiceTitle) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetailsServiceTitle struct{}"
	}

	return strings.Join([]string{"TaskDetailsServiceTitle", string(data)}, " ")
}
