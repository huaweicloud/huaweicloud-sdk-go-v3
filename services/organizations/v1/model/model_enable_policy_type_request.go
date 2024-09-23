package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnablePolicyTypeRequest Request Object
type EnablePolicyTypeRequest struct {

	// 如果正在使用临时安全凭据，则此header是必需的，该值是临时安全凭据的安全令牌（会话令牌）。
	XSecurityToken *string `json:"X-Security-Token,omitempty"`

	Body *PolicyTypeReqBody `json:"body,omitempty"`
}

func (o EnablePolicyTypeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnablePolicyTypeRequest struct{}"
	}

	return strings.Join([]string{"EnablePolicyTypeRequest", string(data)}, " ")
}
