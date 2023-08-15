package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartGraph2Request Request Object
type StartGraph2Request struct {

	// 图ID。
	GraphId string `json:"graph_id"`

	Body *StartGraphReq `json:"body,omitempty"`
}

func (o StartGraph2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartGraph2Request struct{}"
	}

	return strings.Join([]string{"StartGraph2Request", string(data)}, " ")
}
