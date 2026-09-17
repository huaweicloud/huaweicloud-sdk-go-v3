package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppConfigsTemplateResponse Response Object
type DeleteAppConfigsTemplateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAppConfigsTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppConfigsTemplateResponse struct{}"
	}

	return strings.Join([]string{"DeleteAppConfigsTemplateResponse", string(data)}, " ")
}
