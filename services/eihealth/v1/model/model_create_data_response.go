package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateDataResponse struct {

	// 创建文件夹返回体
	Id             *string `json:"id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDataResponse struct{}"
	}

	return strings.Join([]string{"CreateDataResponse", string(data)}, " ")
}
