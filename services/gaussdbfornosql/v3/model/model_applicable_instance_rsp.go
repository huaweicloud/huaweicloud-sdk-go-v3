package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApplicableInstanceRsp struct {

	// 实例ID。
	Id string `json:"id"`

	// 实例名称。
	Name string `json:"name"`
}

func (o ApplicableInstanceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplicableInstanceRsp struct{}"
	}

	return strings.Join([]string{"ApplicableInstanceRsp", string(data)}, " ")
}
