package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePortResponseModel 通道号删除返回体。
type DeletePortResponseModel struct {
	Data *DeletePortResponseModelData `json:"data,omitempty"`
}

func (o DeletePortResponseModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePortResponseModel struct{}"
	}

	return strings.Join([]string{"DeletePortResponseModel", string(data)}, " ")
}
