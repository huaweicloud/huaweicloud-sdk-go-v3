package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFeaturesResponse Response Object
type ListFeaturesResponse struct {

	// 配置列表。
	Configs *[]Feature `json:"configs,omitempty"`

	// 配置项总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListFeaturesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFeaturesResponse struct{}"
	}

	return strings.Join([]string{"ListFeaturesResponse", string(data)}, " ")
}
