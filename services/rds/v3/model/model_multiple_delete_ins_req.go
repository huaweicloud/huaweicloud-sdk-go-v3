package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MultipleDeleteInsReq struct {

	// 实例id列表
	InstanceIds []string `json:"instance_ids"`
}

func (o MultipleDeleteInsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MultipleDeleteInsReq struct{}"
	}

	return strings.Join([]string{"MultipleDeleteInsReq", string(data)}, " ")
}
