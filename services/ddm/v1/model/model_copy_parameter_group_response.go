package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CopyParameterGroupResponse Response Object
type CopyParameterGroupResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CopyParameterGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CopyParameterGroupResponse struct{}"
	}

	return strings.Join([]string{"CopyParameterGroupResponse", string(data)}, " ")
}
