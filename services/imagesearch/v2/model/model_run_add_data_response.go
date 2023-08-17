package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunAddDataResponse Response Object
type RunAddDataResponse struct {

	// 添加数据完成返回success。
	Result *string `json:"result,omitempty"`

	Data           *AddDataRestInfo `json:"data,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o RunAddDataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunAddDataResponse struct{}"
	}

	return strings.Join([]string{"RunAddDataResponse", string(data)}, " ")
}
