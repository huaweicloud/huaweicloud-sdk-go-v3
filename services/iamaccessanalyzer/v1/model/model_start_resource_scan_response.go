package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartResourceScanResponse Response Object
type StartResourceScanResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o StartResourceScanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartResourceScanResponse struct{}"
	}

	return strings.Join([]string{"StartResourceScanResponse", string(data)}, " ")
}
