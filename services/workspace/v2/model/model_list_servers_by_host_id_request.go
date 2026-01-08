package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersByHostIdRequest Request Object
type ListServersByHostIdRequest struct {

	// 云办公主机id。
	HostId string `json:"host_id"`
}

func (o ListServersByHostIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersByHostIdRequest struct{}"
	}

	return strings.Join([]string{"ListServersByHostIdRequest", string(data)}, " ")
}
