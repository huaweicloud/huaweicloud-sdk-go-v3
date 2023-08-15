package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckRepositoryDuplicateNameResponse Response Object
type CheckRepositoryDuplicateNameResponse struct {

	// 重名校验是否通过,true:校验通过不重名,false:校验不通过重名
	Result         *bool `json:"result,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckRepositoryDuplicateNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRepositoryDuplicateNameResponse struct{}"
	}

	return strings.Join([]string{"CheckRepositoryDuplicateNameResponse", string(data)}, " ")
}
