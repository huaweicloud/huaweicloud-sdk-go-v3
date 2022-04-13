package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateKmsTagRequestBody struct {
	Tag *TagItem `json:"tag,omitempty"`
}

func (o CreateKmsTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKmsTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateKmsTagRequestBody", string(data)}, " ")
}
