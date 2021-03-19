package model

import (
	"encoding/json"

	"strings"
)

type SignBindingReq struct {
	// 签名密钥编号

	SignId string `json:"sign_id"`
	// API的发布记录编号

	PublishIds []string `json:"publish_ids"`
}

func (o SignBindingReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SignBindingReq struct{}"
	}

	return strings.Join([]string{"SignBindingReq", string(data)}, " ")
}
