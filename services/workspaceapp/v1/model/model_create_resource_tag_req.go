package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateResourceTagReq tags
type CreateResourceTagReq struct {

	// 标签列表。
	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o CreateResourceTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateResourceTagReq struct{}"
	}

	return strings.Join([]string{"CreateResourceTagReq", string(data)}, " ")
}
