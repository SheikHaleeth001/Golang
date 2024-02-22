package controller

import (
	"fmt"
	"httpserver/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequest(c *gin.Context) {
	var requestBody entity.HttpRequestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Print the received request
	fmt.Printf("Received request: %+v\n", requestBody)

	// Create attributes section
	attributes := map[string]interface{}{
		"form_varient": map[string]interface{}{
			"value": requestBody.Atrv1,
			"type":  requestBody.Atrt1,
		},
		"ref": map[string]interface{}{
			"value": requestBody.Atrv2,
			"type":  requestBody.Atrt2,
		},
	}

	// Create traits section
	traits := map[string]interface{}{
		"name": map[string]interface{}{
			"value": requestBody.Uatrv1,
			"type":  requestBody.Uatrt1,
		},
		"email": map[string]interface{}{
			"value": requestBody.Uatrv2,
			"type":  requestBody.Uatrt2,
		},
		"age": map[string]interface{}{
			"value": requestBody.Uatrv3,
			"type":  requestBody.Uatrt3,
		},
	}

	// Respond to the client
	jsonResponse := entity.JsonResponse{
		Message:         "Request received successfully",
		Event:           requestBody.Ev,
		EventType:       requestBody.Et,
		AppId:           requestBody.Id,
		UserId:          requestBody.Uid,
		MessageId:       requestBody.Mid,
		PageTitle:       requestBody.T,
		PageURL:         requestBody.P,
		BrowserLanguage: requestBody.L,
		ScreenSize:      requestBody.Sc,
		Attributes:      attributes,
		Traits:          traits,
	}

	c.IndentedJSON(http.StatusOK, jsonResponse)
}

func ChannalRequest(c *gin.Context) {
	var request entity.HttpRequestBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attributes := make(map[string]interface{})
	if request.Atrv1 != "" {
		attributes["button_text"] = map[string]interface{}{"value": request.Atrv1, "type": request.Atrt1}
	}
	if request.Atrv2 != "" {
		attributes["color_variation"] = map[string]interface{}{"value": request.Atrv2, "type": request.Atrt2}
	}

	traits := make(map[string]interface{})
	if request.Uatrv1 != "" {
		traits["user_score"] = map[string]interface{}{"value": request.Uatrv1, "type": request.Uatrt1}
	}
	if request.Uatrv2 != "" {
		traits["gender"] = map[string]interface{}{"value": request.Uatrv2, "type": request.Uatrt2}
	}
	if request.Uatrv3 != "" {
		traits["tracking_code"] = map[string]interface{}{"value": request.Uatrv3, "type": request.Uatrt3}
	}

	jsonResponse := entity.JsonResponse{
		Message:         "Request received successfully",
		Event:           request.Ev,
		EventType:       request.Et,
		AppId:           request.Id,
		UserId:          request.Uid,
		MessageId:       request.Mid,
		PageTitle:       request.T,
		PageURL:         request.P,
		BrowserLanguage: request.L,
		ScreenSize:      request.Sc,
		Attributes:      attributes,
		Traits:          traits,
	}

	fmt.Println("Generated JSON response:", jsonResponse)
	c.IndentedJSON(http.StatusOK, jsonResponse)

}

func Worker(c *gin.Context) {
	var request entity.WorkerRequestBody
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	attributes := make(map[string]interface{})
	if request.Atrv1 != "" {
		attributes["button_text"] = map[string]interface{}{"value": request.Atrv1, "type": request.Atrt1}
	}
	if request.Atrv2 != "" {
		attributes["color_variation"] = map[string]interface{}{"value": request.Atrv2, "type": request.Atrt2}
	}
	if request.Atrv3 != "" {
		attributes["page_path"] = map[string]interface{}{"value": request.Atrv3, "type": request.Atrt3}
	}
	if request.Atrv4 != "" {
		attributes["source"] = map[string]interface{}{"value": request.Atrv4, "type": request.Atrt4}
	}

	traits := make(map[string]interface{})
	if request.Uatrv1 != "" {
		traits["user_score"] = map[string]interface{}{"value": request.Uatrv1, "type": request.Uatrt1}
	}
	if request.Uatrv2 != "" {
		traits["gender"] = map[string]interface{}{"value": request.Uatrv2, "type": request.Uatrt2}
	}
	if request.Uatrv3 != "" {
		traits["tracking_code"] = map[string]interface{}{"value": request.Uatrv3, "type": request.Uatrt3}
	}
	if request.Uatrv4 != "" {
		traits["phone"] = map[string]interface{}{"value": request.Uatrv4, "type": request.Uatrt4}
	}
	if request.Uatrv5 != "" {
		traits["coupon_clicked"] = map[string]interface{}{"value": request.Uatrv5, "type": request.Uatrt5}
	}
	if request.Uatrv6 != "" {
		traits["opt_out"] = map[string]interface{}{"value": request.Uatrv6, "type": request.Uatrt6}
	}

	jsonResponse := entity.JsonResponse{
		Message:         "Request received successfully",
		Event:           request.Ev,
		EventType:       request.Et,
		AppId:           request.Id,
		UserId:          request.Uid,
		MessageId:       request.Mid,
		PageTitle:       request.T,
		PageURL:         request.P,
		BrowserLanguage: request.L,
		ScreenSize:      request.Sc,
		Attributes:      attributes,
		Traits:          traits,
	}

	fmt.Println("Generated JSON response:", jsonResponse)
	c.IndentedJSON(http.StatusOK, jsonResponse)
}
