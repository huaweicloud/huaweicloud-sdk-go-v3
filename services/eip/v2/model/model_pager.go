/*
    * EIP
    *
    * 云服务接口
    *
*/

package model

// marker分页结构
type Pager struct {
	// 页码url
	Href string `json:"href,omitempty"`
	// next:下一页  previous:前一页
	Rel string `json:"rel,omitempty"`
}
