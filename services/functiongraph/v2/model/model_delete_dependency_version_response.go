package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteDependencyVersionResponse Response Object
type DeleteDependencyVersionResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteDependencyVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteDependencyVersionResponse struct{}"
	}

	return strings.Join([]string{"DeleteDependencyVersionResponse", string(data)}, " ")
}
