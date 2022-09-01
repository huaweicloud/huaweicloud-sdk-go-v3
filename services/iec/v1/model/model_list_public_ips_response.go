package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPublicIpsResponse struct {

	// 弹性公网IP数目。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 弹性公网IP数组对象。
	Publicips      *[]PublicIp `json:"publicips,omitempty" xml:"publicips"`
	HttpStatusCode int         `json:"-"`
}

func (o ListPublicIpsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublicIpsResponse struct{}"
	}

	return strings.Join([]string{"ListPublicIpsResponse", string(data)}, " ")
}
