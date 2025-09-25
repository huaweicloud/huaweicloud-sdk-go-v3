package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteDaemonsetRequestBody 批量卸载请求体
type BatchDeleteDaemonsetRequestBody struct {

	// 批量卸载列表
	ClusterIds []string `json:"cluster_ids"`

	// 调用服务，标识cce免费体检报告，cce调用传参为cce:   - hss hss服务   - cce cce服务
	InvokedService *string `json:"invoked_service,omitempty"`
}

func (o BatchDeleteDaemonsetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteDaemonsetRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteDaemonsetRequestBody", string(data)}, " ")
}
