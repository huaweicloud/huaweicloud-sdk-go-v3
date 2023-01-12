package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UploadFromObs2Response struct {

	// 元数据的id。
	Id *string `json:"id,omitempty"`

	// 元数据的名字。
	Name           *string `json:"name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UploadFromObs2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UploadFromObs2Response struct{}"
	}

	return strings.Join([]string{"UploadFromObs2Response", string(data)}, " ")
}
