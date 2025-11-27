package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UcsConstraintTemplateTagRequest struct {

	// 标签，多个标签使用‘;’分隔
	Tag *string `json:"tag,omitempty"`
}

func (o UcsConstraintTemplateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UcsConstraintTemplateTagRequest struct{}"
	}

	return strings.Join([]string{"UcsConstraintTemplateTagRequest", string(data)}, " ")
}
