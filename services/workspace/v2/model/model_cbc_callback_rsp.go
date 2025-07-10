package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CbcCallbackRsp CBC回调创建包周期桌面时的响应体。注意：根据云运营平台的API规范，部分参数为驼峰型，不能修改为下划线连接，API规范检查时需要忽略。
type CbcCallbackRsp struct {
}

func (o CbcCallbackRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CbcCallbackRsp struct{}"
	}

	return strings.Join([]string{"CbcCallbackRsp", string(data)}, " ")
}
