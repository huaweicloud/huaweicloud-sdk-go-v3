package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowModelServiceListResponse Response Object
type ShowModelServiceListResponse struct {

	// 推理服务总数
	Total *int32 `json:"total,omitempty"`

	// 模型服务列表
	Services       *[]ModelServiceRsp `json:"services,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ShowModelServiceListResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowModelServiceListResponse struct{}"
	}

	return strings.Join([]string{"ShowModelServiceListResponse", string(data)}, " ")
}
