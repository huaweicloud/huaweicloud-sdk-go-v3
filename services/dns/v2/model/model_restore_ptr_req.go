package model

import (
	"encoding/json"

	"strings"
)

type RestorePtrReq struct {
	// PTR记录对应的域名。  此处值为null。
	Ptrdname string `json:"ptrdname"`
}

func (o RestorePtrReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestorePtrReq struct{}"
	}

	return strings.Join([]string{"RestorePtrReq", string(data)}, " ")
}
