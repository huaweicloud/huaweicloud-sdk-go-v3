package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAgenciesTaskResponse Response Object
type CreateAgenciesTaskResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o CreateAgenciesTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAgenciesTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateAgenciesTaskResponse", string(data)}, " ")
}
