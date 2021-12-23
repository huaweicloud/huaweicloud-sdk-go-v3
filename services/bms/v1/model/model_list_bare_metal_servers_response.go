package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListBareMetalServersResponse struct {
	// 裸金属服务器详情列表

	Servers *[]ServerDetails `json:"servers,omitempty"`
	// 裸金属服务器的列表总数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListBareMetalServersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBareMetalServersResponse struct{}"
	}

	return strings.Join([]string{"ListBareMetalServersResponse", string(data)}, " ")
}
