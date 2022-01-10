package v3

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/vss/v3/model"
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

//认证租户的域名资产
func (c *VssClient) AuthorizeDomains(request *model.AuthorizeDomainsRequest) (*model.AuthorizeDomainsResponse, error) {
	requestDef := GenReqDefForAuthorizeDomains()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.AuthorizeDomainsResponse), nil
	}
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

//获取域名漏洞扫描结果
func (c *VssClient) ShowResults(request *model.ShowResultsRequest) (*model.ShowResultsResponse, error) {
	requestDef := GenReqDefForShowResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowResultsResponse), nil
	}
}

//创建域名漏洞扫描任务并启动
func (c *VssClient) CreateTasks(request *model.CreateTasksRequest) (*model.CreateTasksResponse, error) {
	requestDef := GenReqDefForCreateTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateTasksResponse), nil
	}
}

//获取域名漏洞扫描任务详情
func (c *VssClient) ShowTasks(request *model.ShowTasksRequest) (*model.ShowTasksResponse, error) {
	requestDef := GenReqDefForShowTasks()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowTasksResponse), nil
	}
}
