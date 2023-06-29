package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartTestDatasourceResponse Response Object
type StartTestDatasourceResponse struct {

	// 返回结果 - true (成功)
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o StartTestDatasourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartTestDatasourceResponse struct{}"
	}

	return strings.Join([]string{"StartTestDatasourceResponse", string(data)}, " ")
}
