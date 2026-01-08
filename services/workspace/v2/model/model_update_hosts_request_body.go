package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHostsRequestBody 更新云办公主机的请求体。
type UpdateHostsRequestBody struct {

	// 云办公主机。
	Hosts []UpdateHostParam `json:"hosts"`
}

func (o UpdateHostsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHostsRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateHostsRequestBody", string(data)}, " ")
}
