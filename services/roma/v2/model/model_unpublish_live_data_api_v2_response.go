package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnpublishLiveDataApiV2Response Response Object
type UnpublishLiveDataApiV2Response struct {
	HttpStatusCode int `json:"-"`
}

func (o UnpublishLiveDataApiV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnpublishLiveDataApiV2Response struct{}"
	}

	return strings.Join([]string{"UnpublishLiveDataApiV2Response", string(data)}, " ")
}
