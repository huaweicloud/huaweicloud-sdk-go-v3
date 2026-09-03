package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDownloadUrlReq 获取技能包下载地址请求。
type CreateDownloadUrlReq struct {

	// 指定下载 region。
	Region string `json:"region"`
}

func (o CreateDownloadUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDownloadUrlReq struct{}"
	}

	return strings.Join([]string{"CreateDownloadUrlReq", string(data)}, " ")
}
