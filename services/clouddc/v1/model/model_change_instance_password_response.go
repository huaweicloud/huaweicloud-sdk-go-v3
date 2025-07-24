package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeInstancePasswordResponse Response Object
type ChangeInstancePasswordResponse struct {

	// 实例返回信息
	Instances      *[]ServersResponsePhysicalServers `json:"instances,omitempty"`
	HttpStatusCode int                               `json:"-"`
}

func (o ChangeInstancePasswordResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeInstancePasswordResponse struct{}"
	}

	return strings.Join([]string{"ChangeInstancePasswordResponse", string(data)}, " ")
}
