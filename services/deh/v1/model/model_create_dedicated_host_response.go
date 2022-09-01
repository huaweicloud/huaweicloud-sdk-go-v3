package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDedicatedHostResponse struct {

	// 已分配的专属主机ID数组。租户可以在这些专属主机上创建云服务器。
	DedicatedHostIds *[]string `json:"dedicated_host_ids,omitempty" xml:"dedicated_host_ids"`
	HttpStatusCode   int       `json:"-"`
}

func (o CreateDedicatedHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDedicatedHostResponse struct{}"
	}

	return strings.Join([]string{"CreateDedicatedHostResponse", string(data)}, " ")
}
