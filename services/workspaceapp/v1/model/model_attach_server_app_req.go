package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachServerAppReq 分发软件信息至镜像实例的请求体。
type AttachServerAppReq struct {

	// 云应用仓库软件唯一标识请求列表。
	Items []string `json:"items"`
}

func (o AttachServerAppReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerAppReq struct{}"
	}

	return strings.Join([]string{"AttachServerAppReq", string(data)}, " ")
}
