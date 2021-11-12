package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPublicIpsResponse struct {
	// 弹性公网IP数目。

	Count *int32 `json:"count,omitempty"`
	// 弹性公网IP数组对象。

	Publicips      *[]PublicIp `json:"publicips,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListPublicIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicIpsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicIpsResponse", string(data)}, " ")
}
