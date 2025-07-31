package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPageNoticesResponse Response Object
type ShowPageNoticesResponse struct {

	// 总数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 容器所属集群
	DataList       *[]ShowPageNoticesInfo `json:"data_list,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowPageNoticesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPageNoticesResponse struct{}"
	}

	return strings.Join([]string{"ShowPageNoticesResponse", string(data)}, " ")
}
