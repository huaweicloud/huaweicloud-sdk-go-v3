package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DomainController 域控制器信息。
type DomainController struct {

	// 域控制器IP地址。
	DcIp string `json:"dc_ip"`

	// 域控制器名称。
	DcName string `json:"dc_name"`
}

func (o DomainController) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DomainController struct{}"
	}

	return strings.Join([]string{"DomainController", string(data)}, " ")
}
