package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServersByHostIdResponse Response Object
type ListServersByHostIdResponse struct {

	// 计算机列表。
	Servers        *[]ListServersRspServers `json:"servers,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListServersByHostIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServersByHostIdResponse struct{}"
	}

	return strings.Join([]string{"ListServersByHostIdResponse", string(data)}, " ")
}
