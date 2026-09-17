package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTopTrendResponse Response Object
type ShowTopTrendResponse struct {

	// Top库表数据列表
	TopDataList    *[]TopDataInfo `json:"top_data_list,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowTopTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowTopTrendResponse", string(data)}, " ")
}
