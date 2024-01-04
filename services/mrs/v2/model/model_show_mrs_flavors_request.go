package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMrsFlavorsRequest Request Object
type ShowMrsFlavorsRequest struct {

	// MRS集群版本，不支持多版本查询 ，例如 MRS%203.1.5.1
	VersionName string `json:"version_name"`
}

func (o ShowMrsFlavorsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMrsFlavorsRequest struct{}"
	}

	return strings.Join([]string{"ShowMrsFlavorsRequest", string(data)}, " ")
}
