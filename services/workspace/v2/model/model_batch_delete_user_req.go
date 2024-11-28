package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteUserReq struct {

	// 桌面用户ID列表。
	UserIds []string `json:"user_ids"`
}

func (o BatchDeleteUserReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteUserReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteUserReq", string(data)}, " ")
}
