package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAtlasResponse Response Object
type DeleteAtlasResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteAtlasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAtlasResponse struct{}"
	}

	return strings.Join([]string{"DeleteAtlasResponse", string(data)}, " ")
}
