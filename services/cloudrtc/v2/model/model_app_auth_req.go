package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改app鉴权请求
type AppAuthReq struct {

	// 开启或关闭URL鉴权
	Enable bool `json:"enable"`

	// 有效期，当开启鉴权时必填。  取值范围：[60，2592000]，缺省为300。  单位：秒。
	Expire *int32 `json:"expire,omitempty"`
}

func (o AppAuthReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppAuthReq struct{}"
	}

	return strings.Join([]string{"AppAuthReq", string(data)}, " ")
}
