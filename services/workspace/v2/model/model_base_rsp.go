package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseRsp struct {
}

func (o BaseRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseRsp struct{}"
	}

	return strings.Join([]string{"BaseRsp", string(data)}, " ")
}
