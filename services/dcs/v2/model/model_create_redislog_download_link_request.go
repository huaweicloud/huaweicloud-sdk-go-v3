package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRedislogDownloadLinkRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// 日志的唯一ID，来自于查询运行日志查询接口

	Id string `json:"id"`
}

func (o CreateRedislogDownloadLinkRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRedislogDownloadLinkRequest struct{}"
	}

	return strings.Join([]string{"CreateRedislogDownloadLinkRequest", string(data)}, " ")
}
