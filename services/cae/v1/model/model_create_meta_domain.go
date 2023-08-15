package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateMetaDomain struct {

	// 域名名称。
	Name string `json:"name"`
}

func (o CreateMetaDomain) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMetaDomain struct{}"
	}

	return strings.Join([]string{"CreateMetaDomain", string(data)}, " ")
}
