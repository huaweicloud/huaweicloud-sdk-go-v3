package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHostRequest Request Object
type DeleteHostRequest struct {

	// 主机资产id
	HostId string `json:"host_id"`
}

func (o DeleteHostRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHostRequest struct{}"
	}

	return strings.Join([]string{"DeleteHostRequest", string(data)}, " ")
}
