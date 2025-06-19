package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportIpBlacklistRequest Request Object
type ImportIpBlacklistRequest struct {

	// 防火墙ID，可通过[防火墙ID获取方式](cfw_02_0028.xml)获取
	FwInstanceId string `json:"fw_instance_id"`

	Body *IpBlacklistImportDto `json:"body,omitempty"`
}

func (o ImportIpBlacklistRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportIpBlacklistRequest struct{}"
	}

	return strings.Join([]string{"ImportIpBlacklistRequest", string(data)}, " ")
}
