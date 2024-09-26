package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateGcbTagRequestBody 创建Tag请求体。
type CreateGcbTagRequestBody struct {
	Tag *RequiredTag `json:"tag"`
}

func (o CreateGcbTagRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateGcbTagRequestBody struct{}"
	}

	return strings.Join([]string{"CreateGcbTagRequestBody", string(data)}, " ")
}
