package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InstallInstanceEipResponse Response Object
type InstallInstanceEipResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o InstallInstanceEipResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallInstanceEipResponse struct{}"
	}

	return strings.Join([]string{"InstallInstanceEipResponse", string(data)}, " ")
}
