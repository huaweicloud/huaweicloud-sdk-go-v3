package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationComponentsResponse Response Object
type CreateApplicationComponentsResponse struct {

	// CMDB分配的uuid。
	Data           *string `json:"data,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateApplicationComponentsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationComponentsResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationComponentsResponse", string(data)}, " ")
}
