package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type KeystoneShowServiceRequest struct {
	// 待查询的服务ID。

	ServiceId string `json:"service_id"`
}

func (o KeystoneShowServiceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeystoneShowServiceRequest struct{}"
	}

	return strings.Join([]string{"KeystoneShowServiceRequest", string(data)}, " ")
}
