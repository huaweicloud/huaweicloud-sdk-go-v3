package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiTopNResponse Response Object
type ListApiTopNResponse struct {

	// 调用信息列表
	Statistics     *[]StatisticForCallDetail `json:"statistics,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ListApiTopNResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiTopNResponse struct{}"
	}

	return strings.Join([]string{"ListApiTopNResponse", string(data)}, " ")
}
