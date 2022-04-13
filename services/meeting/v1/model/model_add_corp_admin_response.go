package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type AddCorpAdminResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o AddCorpAdminResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddCorpAdminResponse struct{}"
	}

	return strings.Join([]string{"AddCorpAdminResponse", string(data)}, " ")
}
