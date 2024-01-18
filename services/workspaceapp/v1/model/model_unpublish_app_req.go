package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishAppReq 删除应用。
type UnpublishAppReq struct {

	// 应用ID列表,单次最多允许操作50个应用。
	Ids []string `json:"ids"`
}

func (o UnpublishAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishAppReq struct{}"
	}

	return strings.Join([]string{"UnpublishAppReq", string(data)}, " ")
}
