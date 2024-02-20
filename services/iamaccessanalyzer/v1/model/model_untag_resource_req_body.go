package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UntagResourceReqBody struct {
	TagKeys []string `json:"tag_keys"`
}

func (o UntagResourceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UntagResourceReqBody struct{}"
	}

	return strings.Join([]string{"UntagResourceReqBody", string(data)}, " ")
}
