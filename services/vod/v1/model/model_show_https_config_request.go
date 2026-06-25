package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpsConfigRequest Request Object
type ShowHttpsConfigRequest struct {

	// 加速域名
	Domain string `json:"domain"`
}

func (o ShowHttpsConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpsConfigRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpsConfigRequest", string(data)}, " ")
}
