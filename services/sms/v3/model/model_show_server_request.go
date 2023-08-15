package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerRequest Request Object
type ShowServerRequest struct {

	// 源端服务器在主机迁移服务中的ID
	SourceId string `json:"source_id"`
}

func (o ShowServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerRequest struct{}"
	}

	return strings.Join([]string{"ShowServerRequest", string(data)}, " ")
}
