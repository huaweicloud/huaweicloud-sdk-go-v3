package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartConnectivityTestResponse Response Object
type StartConnectivityTestResponse struct {

	// 连通性测试结果。
	Result         *[]Result `json:"result,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o StartConnectivityTestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartConnectivityTestResponse struct{}"
	}

	return strings.Join([]string{"StartConnectivityTestResponse", string(data)}, " ")
}
