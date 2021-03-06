package rds

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

func (client *Client) DescribeReplicaUsage(request *DescribeReplicaUsageRequest) (response *DescribeReplicaUsageResponse, err error) {
	response = CreateDescribeReplicaUsageResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeReplicaUsageWithChan(request *DescribeReplicaUsageRequest) (<-chan *DescribeReplicaUsageResponse, <-chan error) {
	responseChan := make(chan *DescribeReplicaUsageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReplicaUsage(request)
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

func (client *Client) DescribeReplicaUsageWithCallback(request *DescribeReplicaUsageRequest, callback func(response *DescribeReplicaUsageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReplicaUsageResponse
		var err error
		defer close(result)
		response, err = client.DescribeReplicaUsage(request)
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

type DescribeReplicaUsageRequest struct {
	*requests.RpcRequest
	ResourceOwnerAccount string `position:"Query" name:"ResourceOwnerAccount"`
	SourceDBInstanceId   string `position:"Query" name:"SourceDBInstanceId"`
	ReplicaId            string `position:"Query" name:"ReplicaId"`
	ResourceOwnerId      string `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string `position:"Query" name:"OwnerAccount"`
	OwnerId              string `position:"Query" name:"OwnerId"`
	SecurityToken        string `position:"Query" name:"SecurityToken"`
}

type DescribeReplicaUsageResponse struct {
	*responses.BaseResponse
	RequestId       string `json:"RequestId" xml:"RequestId"`
	StartTime       string `json:"StartTime" xml:"StartTime"`
	EndTime         string `json:"EndTime" xml:"EndTime"`
	ReplicaId       string `json:"ReplicaId" xml:"ReplicaId"`
	PerformanceKeys struct {
		PerformanceKey []struct {
			Key               string `json:"Key" xml:"Key"`
			Unit              string `json:"Unit" xml:"Unit"`
			ValueFormat       string `json:"ValueFormat" xml:"ValueFormat"`
			PerformanceValues struct {
				PerformanceValue []struct {
					Value string `json:"Value" xml:"Value"`
					Date  string `json:"Date" xml:"Date"`
				} `json:"PerformanceValue" xml:"PerformanceValue"`
			} `json:"PerformanceValues" xml:"PerformanceValues"`
		} `json:"PerformanceKey" xml:"PerformanceKey"`
	} `json:"PerformanceKeys" xml:"PerformanceKeys"`
}

func CreateDescribeReplicaUsageRequest() (request *DescribeReplicaUsageRequest) {
	request = &DescribeReplicaUsageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicaUsage", "", "")
	return
}

func CreateDescribeReplicaUsageResponse() (response *DescribeReplicaUsageResponse) {
	response = &DescribeReplicaUsageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
