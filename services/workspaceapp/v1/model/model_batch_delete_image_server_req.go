package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteImageServerReq 批量删除镜像实例请求。 * 忽略不存在的镜像实例，响应正常。 * 不允许操作状态为 `创建中`、`镜像创建中`的实例，响应异常。 * 不支持资源关联发生变化后，请求删除镜像实例关联资源，任务响应异常。
type BatchDeleteImageServerReq struct {

	// 批量唯一标识请求列表，一次请求数量区间 [1, 20]
	Items []string `json:"items"`

	// 是否同时删除镜像实例关联资源： **⚠ 警告: 关联资源删除，对应的应用将不可用** * `true` 同时删除关联资源，包括APS服务器组，APS服务器，应用组相关资源。镜像产物相关信息保留。 * `false` 只删除镜像实例记录，保留关联资源。
	Recursive *bool `json:"recursive,omitempty"`
}

func (o BatchDeleteImageServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteImageServerReq struct{}"
	}

	return strings.Join([]string{"BatchDeleteImageServerReq", string(data)}, " ")
}
