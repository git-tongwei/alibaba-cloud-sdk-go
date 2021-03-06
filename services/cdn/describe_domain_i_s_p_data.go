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

func (client *Client) DescribeDomainISPData(request *DescribeDomainISPDataRequest) (response *DescribeDomainISPDataResponse, err error) {
	response = CreateDescribeDomainISPDataResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainISPDataWithChan(request *DescribeDomainISPDataRequest) (<-chan *DescribeDomainISPDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainISPDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainISPData(request)
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

func (client *Client) DescribeDomainISPDataWithCallback(request *DescribeDomainISPDataRequest, callback func(response *DescribeDomainISPDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainISPDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainISPData(request)
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

type DescribeDomainISPDataRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Action        string `position:"Query" name:"Action"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
}

type DescribeDomainISPDataResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId"`
	DomainName   string `json:"DomainName"`
	DataInterval string `json:"DataInterval"`
	StartTime    string `json:"StartTime"`
	EndTime      string `json:"EndTime"`
	Value        []struct {
		ISP             string `json:"ISP"`
		Proportion      string `json:"Proportion"`
		IspEname        string `json:"IspEname"`
		AvgObjectSize   string `json:"AvgObjectSize"`
		AvgResponseTime string `json:"AvgResponseTime"`
		Bps             string `json:"Bps"`
		ByteHitRate     string `json:"ByteHitRate"`
		Qps             string `json:"Qps"`
		ReqErrRate      string `json:"ReqErrRate"`
		ReqHitRate      string `json:"ReqHitRate"`
		AvgResponseRate string `json:"AvgResponseRate"`
		TotalBytes      string `json:"TotalBytes"`
		BytesProportion string `json:"BytesProportion"`
		TotalQuery      string `json:"TotalQuery"`
	} `json:"Value"`
}

func CreateDescribeDomainISPDataRequest() (request *DescribeDomainISPDataRequest) {
	request = &DescribeDomainISPDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainISPData", "", "")
	return
}

func CreateDescribeDomainISPDataResponse() (response *DescribeDomainISPDataResponse) {
	response = &DescribeDomainISPDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
