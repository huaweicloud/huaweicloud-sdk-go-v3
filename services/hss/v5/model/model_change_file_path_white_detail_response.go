package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeFilePathWhiteDetailResponse Response Object
type ChangeFilePathWhiteDetailResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ChangeFilePathWhiteDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeFilePathWhiteDetailResponse struct{}"
	}

	return strings.Join([]string{"ChangeFilePathWhiteDetailResponse", string(data)}, " ")
}
