package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateAppReq 批量按字段更新,仅仅更新非null字段。
type BatchUpdateAppReq struct {

	// 应用ID列表,单次最多允许操作20条。
	Ids []string `json:"ids"`
}

func (o BatchUpdateAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateAppReq struct{}"
	}

	return strings.Join([]string{"BatchUpdateAppReq", string(data)}, " ")
}
