package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVhostBody struct {

	// Vhost名称
	Name string `json:"name"`
}

func (o CreateVhostBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVhostBody struct{}"
	}

	return strings.Join([]string{"CreateVhostBody", string(data)}, " ")
}
