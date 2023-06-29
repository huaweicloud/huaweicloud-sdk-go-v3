package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSubjectResponse Response Object
type UpdateSubjectResponse struct {

	// 返回的数据信息
	Data           *interface{} `json:"data,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateSubjectResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSubjectResponse struct{}"
	}

	return strings.Join([]string{"UpdateSubjectResponse", string(data)}, " ")
}
