package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddCorpAdminResponse Response Object
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
