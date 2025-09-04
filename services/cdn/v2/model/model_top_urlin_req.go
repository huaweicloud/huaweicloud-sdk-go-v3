package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopUrlinReq top_url配置
type TopUrlinReq struct {

	// 配置开关
	Enable *bool `json:"enable,omitempty"`
}

func (o TopUrlinReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopUrlinReq struct{}"
	}

	return strings.Join([]string{"TopUrlinReq", string(data)}, " ")
}
