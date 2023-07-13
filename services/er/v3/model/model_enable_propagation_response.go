package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnablePropagationResponse Response Object
type EnablePropagationResponse struct {
	Propagation *Propagation `json:"propagation,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	XClientToken   *string `json:"X-Client-Token,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EnablePropagationResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnablePropagationResponse struct{}"
	}

	return strings.Join([]string{"EnablePropagationResponse", string(data)}, " ")
}
