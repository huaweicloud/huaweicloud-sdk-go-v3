package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchModifyGroupDescriptionReq struct {

	// 消息组列表
	Groups *[]GroupCreateReq `json:"groups,omitempty"`
}

func (o BatchModifyGroupDescriptionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchModifyGroupDescriptionReq struct{}"
	}

	return strings.Join([]string{"BatchModifyGroupDescriptionReq", string(data)}, " ")
}
