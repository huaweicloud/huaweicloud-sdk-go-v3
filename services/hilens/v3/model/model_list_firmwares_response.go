package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFirmwaresResponse Response Object
type ListFirmwaresResponse struct {

	// 固件数
	Count *int32 `json:"count,omitempty"`

	// 固件列表
	Firmwares      *[]ListFirmwaresResponseData `json:"firmwares,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListFirmwaresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFirmwaresResponse struct{}"
	}

	return strings.Join([]string{"ListFirmwaresResponse", string(data)}, " ")
}
