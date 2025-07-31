package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateDaemonsetRequestBody 批量升级请求体
type BatchUpdateDaemonsetRequestBody struct {

	// 批量升级列表
	DataList []UpdateDaemonsetInfo `json:"data_list"`
}

func (o BatchUpdateDaemonsetRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateDaemonsetRequestBody struct{}"
	}

	return strings.Join([]string{"BatchUpdateDaemonsetRequestBody", string(data)}, " ")
}
