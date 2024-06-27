package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateChInstanceRequestBodyHa 部署信息。
type CreateChInstanceRequestBodyHa struct {

	// 部署模式。 取值范围： - Single：单机 - Ha：主备
	Mode string `json:"mode"`
}

func (o CreateChInstanceRequestBodyHa) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChInstanceRequestBodyHa struct{}"
	}

	return strings.Join([]string{"CreateChInstanceRequestBodyHa", string(data)}, " ")
}
