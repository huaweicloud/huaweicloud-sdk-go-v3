package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StackSetVarsUriContentPrimitiveTypeHolder struct {

	// vars_uri对应的文件内容
	VarsUriContent *string `json:"vars_uri_content,omitempty"`
}

func (o StackSetVarsUriContentPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetVarsUriContentPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetVarsUriContentPrimitiveTypeHolder", string(data)}, " ")
}
