package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagResponse Response Object
type CreateTagResponse struct {

	// 标签名称
	Name *string `json:"name,omitempty"`

	// 标签描述
	Message *string `json:"message,omitempty"`

	// 基于commitid
	Target *string `json:"target,omitempty"`

	Commit         *CommitDto `json:"commit,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o CreateTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagResponse struct{}"
	}

	return strings.Join([]string{"CreateTagResponse", string(data)}, " ")
}
