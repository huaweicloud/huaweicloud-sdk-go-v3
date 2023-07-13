package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VarsUriContentPrimitiveTypeHolder struct {

	// vars_uri对应的文件内容
	VarsUriContent *string `json:"vars_uri_content,omitempty"`
}

func (o VarsUriContentPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VarsUriContentPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"VarsUriContentPrimitiveTypeHolder", string(data)}, " ")
}
