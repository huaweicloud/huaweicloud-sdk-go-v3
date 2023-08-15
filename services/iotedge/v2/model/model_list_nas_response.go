package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNasResponse Response Object
type ListNasResponse struct {

	// 北向NA列表
	Nas            *[]QueryNaBriefResponseDto `json:"nas,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ListNasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNasResponse struct{}"
	}

	return strings.Join([]string{"ListNasResponse", string(data)}, " ")
}
