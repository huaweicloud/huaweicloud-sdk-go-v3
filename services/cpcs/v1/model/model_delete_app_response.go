package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppResponse Response Object
type DeleteAppResponse struct {

	// 应用ID
	AppId          *string `json:"app_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteAppResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppResponse", string(data)}, " ")
}
