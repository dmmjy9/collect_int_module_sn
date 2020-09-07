package controller

import "github.com/gin-gonic/gin"

func SendTask(c *gin.Context) {
	availableVendorSlice := []string{
		"h3c",
		"ruijie",
		"huawei",
		"juniper",
		"cisco",
		"arista",
	}
	availableVendor := make(map[string]int)

	for _, val := range availableVendorSlice {
		availableVendor[val] = 1
	}

	postVendor := c.PostForm("vendor")
	postIpaddr := c.PostForm("ipaddr")

	if availableVendor[postVendor] != 1 {
	}
}

func QueryTaskList(c *gin.Context) {

}

func GetResult(c *gin.Context) {

}
