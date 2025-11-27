package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStatisticsResponse Response Object
type UpdateStatisticsResponse struct {

	// 响应结果
	Resp           *string `json:"resp,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStatisticsResponse struct{}"
	}

	return strings.Join([]string{"UpdateStatisticsResponse", string(data)}, " ")
}
