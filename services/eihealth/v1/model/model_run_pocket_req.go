package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPocketReq RunPocket请求体
type RunPocketReq struct {

	// 口袋中心坐标
	Center []float32 `json:"center"`

	// 口袋大小
	Size []float32 `json:"size"`

	// 口袋的padding值
	Padding float32 `json:"padding"`
}

func (o RunPocketReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPocketReq struct{}"
	}

	return strings.Join([]string{"RunPocketReq", string(data)}, " ")
}
