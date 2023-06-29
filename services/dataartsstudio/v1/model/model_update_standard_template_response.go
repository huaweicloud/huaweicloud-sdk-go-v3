package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateStandardTemplateResponse Response Object
type UpdateStandardTemplateResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateStandardTemplateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateStandardTemplateResponse struct{}"
	}

	return strings.Join([]string{"UpdateStandardTemplateResponse", string(data)}, " ")
}
