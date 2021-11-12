package v1

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dsc/v1/model"
)

type DscClient struct {
	HcClient *http_client.HcHttpClient
}

func NewDscClient(hcClient *http_client.HcHttpClient) *DscClient {
	return &DscClient{HcClient: hcClient}
}

func DscClientBuilder() *http_client.HcHttpClientBuilder {
	builder := http_client.NewHcHttpClientBuilder()
	return builder
}

//对数据进行脱敏
func (c *DscClient) BatchAddDataMask(request *model.BatchAddDataMaskRequest) (*model.BatchAddDataMaskResponse, error) {
	requestDef := GenReqDefForBatchAddDataMask()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.BatchAddDataMaskResponse), nil
	}
}

//对json体中数据动态添加水印
func (c *DscClient) CreateDatabaseWaterMark(request *model.CreateDatabaseWaterMarkRequest) (*model.CreateDatabaseWaterMarkResponse, error) {
	requestDef := GenReqDefForCreateDatabaseWaterMark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDatabaseWaterMarkResponse), nil
	}
}

//嵌入文档水印
func (c *DscClient) CreateDocWatermark(request *model.CreateDocWatermarkRequest) (*model.CreateDocWatermarkResponse, error) {
	requestDef := GenReqDefForCreateDocWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateDocWatermarkResponse), nil
	}
}

//给上传的图片添加暗水印，目前支持的图片格式为：*.jpg, *.jpeg, *.jpe, *.png, *.bmp, *.dib, *.rle, *.tiff, *.tif, *.ppm, *.webp, *.tga, *.tpic, *.gif。
func (c *DscClient) CreateImageWatermark(request *model.CreateImageWatermarkRequest) (*model.CreateImageWatermarkResponse, error) {
	requestDef := GenReqDefForCreateImageWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.CreateImageWatermarkResponse), nil
	}
}

//提取请求数据中水印内容
func (c *DscClient) ShowDatabaseWaterMark(request *model.ShowDatabaseWaterMarkRequest) (*model.ShowDatabaseWaterMarkResponse, error) {
	requestDef := GenReqDefForShowDatabaseWaterMark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDatabaseWaterMarkResponse), nil
	}
}

//提取文档水印
func (c *DscClient) ShowDocWatermark(request *model.ShowDocWatermarkRequest) (*model.ShowDocWatermarkResponse, error) {
	requestDef := GenReqDefForShowDocWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowDocWatermarkResponse), nil
	}
}

//提取图片中的暗水印内容
func (c *DscClient) ShowImageWatermark(request *model.ShowImageWatermarkRequest) (*model.ShowImageWatermarkResponse, error) {
	requestDef := GenReqDefForShowImageWatermark()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowImageWatermarkResponse), nil
	}
}

//查询指定任务扫描结果
func (c *DscClient) ShowScanJobResults(request *model.ShowScanJobResultsRequest) (*model.ShowScanJobResultsResponse, error) {
	requestDef := GenReqDefForShowScanJobResults()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScanJobResultsResponse), nil
	}
}

//查询扫描任务列表
func (c *DscClient) ShowScanJobs(request *model.ShowScanJobsRequest) (*model.ShowScanJobsResponse, error) {
	requestDef := GenReqDefForShowScanJobs()

	if resp, err := c.HcClient.Sync(request, requestDef); err != nil {
		return nil, err
	} else {
		return resp.(*model.ShowScanJobsResponse), nil
	}
}
