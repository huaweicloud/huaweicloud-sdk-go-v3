package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OsChangeReq 切换裸金属服务器操作系统接口请求结构体
type OsChangeReq struct {
	OsChange *OsChange `json:"os-change"`
}

func (o OsChangeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OsChangeReq struct{}"
	}

	return strings.Join([]string{"OsChangeReq", string(data)}, " ")
}
