package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPathObjectByIdResponse Response Object
type ShowPathObjectByIdResponse struct {

	// 路径对象
	Paths          *[]LayerPath `json:"paths,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowPathObjectByIdResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPathObjectByIdResponse struct{}"
	}

	return strings.Join([]string{"ShowPathObjectByIdResponse", string(data)}, " ")
}
