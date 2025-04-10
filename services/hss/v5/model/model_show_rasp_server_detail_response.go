package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRaspServerDetailResponse Response Object
type ShowRaspServerDetailResponse struct {

	// 总数
	TotalNum *int64 `json:"total_num,omitempty"`

	// java应用状态列表数据
	DataList       *[]ServerAppStatusResponseInfo `json:"data_list,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ShowRaspServerDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRaspServerDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowRaspServerDetailResponse", string(data)}, " ")
}
