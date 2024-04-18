package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopIoTrafficsResponse Response Object
type ListTopIoTrafficsResponse struct {

	// Top IO列表
	TopIoInfos     *[]TopIoInfo `json:"top_io_infos,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTopIoTrafficsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopIoTrafficsResponse struct{}"
	}

	return strings.Join([]string{"ListTopIoTrafficsResponse", string(data)}, " ")
}
