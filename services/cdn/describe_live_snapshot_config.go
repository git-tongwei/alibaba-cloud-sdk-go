package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeLiveSnapshotConfig(request *DescribeLiveSnapshotConfigRequest) (response *DescribeLiveSnapshotConfigResponse, err error) {
	response = CreateDescribeLiveSnapshotConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveSnapshotConfigWithChan(request *DescribeLiveSnapshotConfigRequest) (<-chan *DescribeLiveSnapshotConfigResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveSnapshotConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveSnapshotConfig(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeLiveSnapshotConfigWithCallback(request *DescribeLiveSnapshotConfigRequest, callback func(response *DescribeLiveSnapshotConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveSnapshotConfigResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveSnapshotConfig(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeLiveSnapshotConfigRequest struct {
	*requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	PageSize      string `position:"Query" name:"PageSize"`
	Action        string `position:"Query" name:"Action"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	PageNum       string `position:"Query" name:"PageNum"`
	StreamName    string `position:"Query" name:"StreamName"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
	Order         string `position:"Query" name:"Order"`
}

type DescribeLiveSnapshotConfigResponse struct {
	*responses.BaseResponse
	RequestId                    string `json:"RequestId"`
	PageNum                      int    `json:"PageNum"`
	PageSize                     int    `json:"PageSize"`
	Order                        string `json:"Order"`
	TotalNum                     int    `json:"TotalNum"`
	TotalPage                    int    `json:"TotalPage"`
	LiveStreamSnapshotConfigList []struct {
		DomainName         string `json:"DomainName"`
		AppName            string `json:"AppName"`
		TimeInterval       int    `json:"TimeInterval"`
		OssEndpoint        string `json:"OssEndpoint"`
		OssBucket          string `json:"OssBucket"`
		OverwriteOssObject string `json:"OverwriteOssObject"`
		SequenceOssObject  string `json:"SequenceOssObject"`
		CreateTime         string `json:"CreateTime"`
	} `json:"LiveStreamSnapshotConfigList"`
}

func CreateDescribeLiveSnapshotConfigRequest() (request *DescribeLiveSnapshotConfigRequest) {
	request = &DescribeLiveSnapshotConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveSnapshotConfig", "", "")
	return
}

func CreateDescribeLiveSnapshotConfigResponse() (response *DescribeLiveSnapshotConfigResponse) {
	response = &DescribeLiveSnapshotConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
