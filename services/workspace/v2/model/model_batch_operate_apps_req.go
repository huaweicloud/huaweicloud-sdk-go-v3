package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchOperateAppsReq 批量操作应用。
type BatchOperateAppsReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 50]。
	Items []string `json:"items"`
}

func (o BatchOperateAppsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchOperateAppsReq struct{}"
	}

	return strings.Join([]string{"BatchOperateAppsReq", string(data)}, " ")
}
