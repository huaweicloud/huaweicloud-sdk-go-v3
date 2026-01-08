package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListResourcePackagesRequest Request Object
type ListResourcePackagesRequest struct {

	// 资源包的资源规格编码。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
}

func (o ListResourcePackagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourcePackagesRequest struct{}"
	}

	return strings.Join([]string{"ListResourcePackagesRequest", string(data)}, " ")
}
