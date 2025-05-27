package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type GetOper struct {
	GetKv *GetKv `bson:"get_kv,omitempty"`
}

func (o GetOper) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetOper struct{}"
	}

	return strings.Join([]string{"GetOper", string(data)}, " ")
}
