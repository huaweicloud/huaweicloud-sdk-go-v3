package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteStorageClaimReq 根据storage_claim_id删除对应的共享存储目录
type DeleteStorageClaimReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]
	Items []string `json:"items"`
}

func (o DeleteStorageClaimReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteStorageClaimReq struct{}"
	}

	return strings.Join([]string{"DeleteStorageClaimReq", string(data)}, " ")
}
