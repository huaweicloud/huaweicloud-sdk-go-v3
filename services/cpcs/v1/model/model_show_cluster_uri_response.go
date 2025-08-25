package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterUriResponse Response Object
type ShowClusterUriResponse struct {

	// 集群业务管理跳转链接
	Uri            *string `json:"uri,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowClusterUriResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterUriResponse struct{}"
	}

	return strings.Join([]string{"ShowClusterUriResponse", string(data)}, " ")
}
