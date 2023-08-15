package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SourceServerByTask 源端服务器信息
type SourceServerByTask struct {

	// 源端服务器ID
	Id string `json:"id"`
}

func (o SourceServerByTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceServerByTask struct{}"
	}

	return strings.Join([]string{"SourceServerByTask", string(data)}, " ")
}
