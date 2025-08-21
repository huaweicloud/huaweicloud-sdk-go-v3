package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAccessDetailResponse Response Object
type ShowAccessDetailResponse struct {
	Data           *AccessDetailVo `json:"data,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ShowAccessDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAccessDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowAccessDetailResponse", string(data)}, " ")
}
