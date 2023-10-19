package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCloudConnectionRequest Request Object
type DeleteCloudConnectionRequest struct {

	// 资源的Id。
	Id string `json:"id"`
}

func (o DeleteCloudConnectionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCloudConnectionRequest struct{}"
	}

	return strings.Join([]string{"DeleteCloudConnectionRequest", string(data)}, " ")
}
