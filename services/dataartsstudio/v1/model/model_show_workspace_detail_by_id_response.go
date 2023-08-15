package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkspaceDetailByIdResponse Response Object
type ShowWorkspaceDetailByIdResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowWorkspaceDetailByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkspaceDetailByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkspaceDetailByIdResponse", string(data)}, " ")
}
