package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterDomainRequest Request Object
type RegisterDomainRequest struct {

	// 桌面ID。
	DesktopId string `json:"desktop_id"`
}

func (o RegisterDomainRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDomainRequest struct{}"
	}

	return strings.Join([]string{"RegisterDomainRequest", string(data)}, " ")
}
