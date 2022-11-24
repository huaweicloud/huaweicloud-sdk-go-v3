package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreatePerformanceResourceReq struct {

	// 存储空间，单位GB
	Space int32 `json:"space"`

	// 购买数量
	Count int32 `json:"count"`
}

func (o CreatePerformanceResourceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePerformanceResourceReq struct{}"
	}

	return strings.Join([]string{"CreatePerformanceResourceReq", string(data)}, " ")
}
