package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Handler 备注信息
type Handler struct {
}

func (o Handler) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Handler struct{}"
	}

	return strings.Join([]string{"Handler", string(data)}, " ")
}
