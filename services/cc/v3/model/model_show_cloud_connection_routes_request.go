package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudConnectionRoutesRequest Request Object
type ShowCloudConnectionRoutesRequest struct {

	// 资源的Id。
	Id string `json:"id"`
}

func (o ShowCloudConnectionRoutesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudConnectionRoutesRequest struct{}"
	}

	return strings.Join([]string{"ShowCloudConnectionRoutesRequest", string(data)}, " ")
}
