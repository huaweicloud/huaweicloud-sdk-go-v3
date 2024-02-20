package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TagResourceReqBody struct {
	Tags []Tag `json:"tags"`
}

func (o TagResourceReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TagResourceReqBody struct{}"
	}

	return strings.Join([]string{"TagResourceReqBody", string(data)}, " ")
}
