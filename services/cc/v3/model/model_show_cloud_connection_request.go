package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudConnectionRequest Request Object
type ShowCloudConnectionRequest struct {

	// 资源的Id。
	Id string `json:"id"`
}

func (o ShowCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"ShowCloudConnectionRequest", string(data)}, " ")
}
