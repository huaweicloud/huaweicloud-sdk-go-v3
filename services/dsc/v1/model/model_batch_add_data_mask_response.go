package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type BatchAddDataMaskResponse struct {

	// 脱敏后的数据的数据列表，结构与请求中结构相同
	MaskedData     *[]map[string]interface{} `json:"masked_data,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o BatchAddDataMaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddDataMaskResponse struct{}"
	}

	return strings.Join([]string{"BatchAddDataMaskResponse", string(data)}, " ")
}
