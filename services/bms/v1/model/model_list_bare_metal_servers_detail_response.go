package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListBareMetalServersDetailResponse Response Object
type ListBareMetalServersDetailResponse struct {

	// 裸金属服务器详情列表
	Servers        *[]ServerListDetails `json:"servers,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListBareMetalServersDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBareMetalServersDetailResponse struct{}"
	}

	return strings.Join([]string{"ListBareMetalServersDetailResponse", string(data)}, " ")
}
