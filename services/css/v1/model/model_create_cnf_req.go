package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCnfReq struct {
	// 名称。

	Name string `json:"name"`
	// 配置文件内容。

	ConfContent string `json:"confContent"`

	Setting *Setting `json:"setting"`
}

func (o CreateCnfReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCnfReq struct{}"
	}

	return strings.Join([]string{"CreateCnfReq", string(data)}, " ")
}
