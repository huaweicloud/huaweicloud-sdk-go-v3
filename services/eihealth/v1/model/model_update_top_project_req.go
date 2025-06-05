package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopProjectReq 置顶和取消置顶空间请求体。
type UpdateTopProjectReq struct {

	// 是否置顶，true表示置顶，false表示取消置顶。
	IsTop bool `json:"is_top"`
}

func (o UpdateTopProjectReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopProjectReq struct{}"
	}

	return strings.Join([]string{"UpdateTopProjectReq", string(data)}, " ")
}
