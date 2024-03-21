package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateV2TagRequestBody 标签对象
type CreateV2TagRequestBody struct {
	Tag *CreateV2TagRequestBodyTag `json:"tag"`
}

func (o CreateV2TagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateV2TagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateV2TagRequestBody", string(data)}, " ")
}
