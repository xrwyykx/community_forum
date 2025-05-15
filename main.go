package main

import (
	"communityforum/global"
	"communityforum/router"

	"github.com/gin-gonic/gin"
)

func main() {
	//加载配置文件
	global.LoadConfig()
	gin.SetMode(gin.ReleaseMode)
	//连接redis
	global.InitRedis()
	//连接数据库
	global.InitDB()
	////启动路由
	router.InitRouterAndStartServer()

}

//func ocrHandwritingWithBaidu(imgURL string) (string, error) {
//	apiKey := "aLnqDRQEwLa7yX7XhJJKa1X7"
//	secretKey := "6IkAaJhaZAGPKmaJADy2EPAdkkUauyVD"
//	tokenURL := fmt.Sprintf("https://aip.baidubce.com/oauth/2.0/token?grant_type=client_credentials&client_id=%s&client_secret=%s", apiKey, secretKey)
//
//	// 1. 获取Access Token
//	tokenResp, _ := http.Get(tokenURL)
//	defer tokenResp.Body.Close()
//	tokenBody, _ := io.ReadAll(tokenResp.Body)
//	var tokenResult map[string]interface{}
//	json.Unmarshal(tokenBody, &tokenResult)
//	accessToken := tokenResult["access_token"].(string)
//
//	// 2. 调用手写OCR接口
//	ocrURL := fmt.Sprintf("https://aip.baidubce.com/rest/2.0/ocr/v1/handwriting?access_token=%s", accessToken)
//	formData := url.Values{"url": {imgURL}, "recognize_granularity": {"big"}} // 大颗粒度适合手写
//
//	resp, err := http.PostForm(ocrURL, formData)
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	// 3. 解析结果（百度返回JSON结构较复杂，需多层解析）
//	ocrBody, _ := io.ReadAll(resp.Body)
//	var ocrResult map[string]interface{}
//	json.Unmarshal(ocrBody, &ocrResult)
//
//	var text strings.Builder
//	for _, word := range ocrResult["words_result"].([]interface{}) {
//		text.WriteString(word.(map[string]interface{})["words"].(string) + "\n")
//	}
//	return text.String(), nil
//}
