package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteResourceTagReq tags
type DeleteResourceTagReq struct {

	// 标签列表
	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o DeleteResourceTagReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteResourceTagReq struct{}"
	}

	return strings.Join([]string{"DeleteResourceTagReq", string(data)}, " ")
}
