package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCollectConfigResponse Response Object
type CreateCollectConfigResponse struct {

	// 保存云服务采集配置响应
	Body           *[]ConfigResponse `json:"body,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o CreateCollectConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCollectConfigResponse struct{}"
	}

	return strings.Join([]string{"CreateCollectConfigResponse", string(data)}, " ")
}
