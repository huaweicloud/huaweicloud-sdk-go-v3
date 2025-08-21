package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTagResponse Response Object
type ShowTagResponse struct {

	// 标签名称
	Name *string `json:"name,omitempty"`

	// 标签描述
	Message *string `json:"message,omitempty"`

	// 基于commitid
	Target *string `json:"target,omitempty"`

	Commit *CommitDto `json:"commit,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowTagResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTagResponse struct{}"
	}

	return strings.Join([]string{"ShowTagResponse", string(data)}, " ")
}
