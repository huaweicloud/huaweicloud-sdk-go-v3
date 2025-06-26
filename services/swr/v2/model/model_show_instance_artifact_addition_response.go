package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceArtifactAdditionResponse Response Object
type ShowInstanceArtifactAdditionResponse struct {

	// 构建历史列表
	BuildHistories *[]BuildHistory `json:"build_histories,omitempty"`

	// 构建历史条数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceArtifactAdditionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceArtifactAdditionResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceArtifactAdditionResponse", string(data)}, " ")
}
