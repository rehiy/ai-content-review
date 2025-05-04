package server

import (
	"encoding/json"
	"net/http"

	"ai-content-review/service"
)

// 处理文本审核的 HTTP POST 请求
// 解码请求负载，处理文本，并以 JSON 格式返回结果
func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "仅允许 POST 方法", http.StatusMethodNotAllowed)
		return
	}

	req := ReviewRequest{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "无效的请求负载", http.StatusBadRequest)
		return
	}

	result, err := service.ReviewText(req.Content)
	if err != nil {
		http.Error(w, "处理请求时出错: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ReviewResponse{Result: result})
}

// 请求的结构体
type ReviewRequest struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

// 响应的结构体
type ReviewResponse struct {
	Result string `json:"result"`
}
