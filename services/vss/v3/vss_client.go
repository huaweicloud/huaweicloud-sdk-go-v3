package v3

import (
	http_client "github.com/RandolphCYG/huaweicloud-sdk-go-v3/core"

	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/services/vss/v3/model"
)

type VssClient struct {
	HcClient *http_client.HcHttpClient
}

func NewVssClient(hcClient *http_client.HcHttpClient) *VssClient {
	return &VssClient{HcClient: hcClient}
}

func VssClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//创建租户的域名资产
func (c *VssClient) CreateDomains(request *model.CreateDomainsRequest) (*model.CreateDomainsResponse, error) {
	requestDef := GenReqDefForCreateDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDomainsResponse), nil
	}
}

//删除租户的域名资产
func (c *VssClient) DeleteDomains(request *model.DeleteDomainsRequest) (*model.DeleteDomainsResponse, error) {
	requestDef := GenReqDefForDeleteDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.DeleteDomainsResponse), nil
	}
}

//获取租户的所有域名资产
func (c *VssClient) ListDomains(request *model.ListDomainsRequest) (*model.ListDomainsResponse, error) {
	requestDef := GenReqDefForListDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ListDomainsResponse), nil
	}
}
