package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListApiVersionResponse Response Object
type ListApiVersionResponse struct {

	// 描述主机迁移服务API版本信息列表。
	Versions       *[]Version `json:"versions,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListApiVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ListApiVersionResponse", string(data)}, " ")
}
