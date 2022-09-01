package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Context struct {

	// 上下文cluster信息。
	Cluster *string `json:"cluster,omitempty" xml:"cluster"`

	// 上下文user信息。
	User *string `json:"user,omitempty" xml:"user"`
}

func (o Context) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Context struct{}"
	}

	return strings.Join([]string{"Context", string(data)}, " ")
}
