package utils

import github.com/igorlev91/golang-grpc-ecommerce/pb/common"

func SuccessResponse() *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    "Success",
	}
}
