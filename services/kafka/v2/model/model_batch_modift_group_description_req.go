package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchModiftGroupDescriptionReq struct {

	// 消费组列表
	Groups *[]GroupCreateReq `json:"groups,omitempty"`
}

func (o BatchModiftGroupDescriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModiftGroupDescriptionReq struct{}"
	}

	return strings.Join([]string{"BatchModiftGroupDescriptionReq", string(data)}, " ")
}
