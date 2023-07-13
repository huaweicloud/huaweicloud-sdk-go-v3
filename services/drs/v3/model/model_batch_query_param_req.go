package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchQueryParamReq
type BatchQueryParamReq struct {

	// 查询任务ID集合。
	Jobs []string `json:"jobs"`

	// 是否重新获取数据库参数，1代表是，0代表否（从缓存中获取），第一次调用时请设置为1。
	Refresh string `json:"refresh"`
}

func (o BatchQueryParamReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchQueryParamReq struct{}"
	}

	return strings.Join([]string{"BatchQueryParamReq", string(data)}, " ")
}
