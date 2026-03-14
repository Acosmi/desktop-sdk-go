package acosmi

import "time"

// ---------- OAuth ----------

// ServerMetadata OAuth Authorization Server 元数据 (RFC 8414)
type ServerMetadata struct {
	Issuer                string   `json:"issuer"`
	AuthorizationEndpoint string   `json:"authorization_endpoint"`
	TokenEndpoint         string   `json:"token_endpoint"`
	RevocationEndpoint    string   `json:"revocation_endpoint"`
	RegistrationEndpoint  string   `json:"registration_endpoint"`
	ScopesSupported       []string `json:"scopes_supported"`
}

// TokenResponse OAuth token 响应
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

// TokenSet 持久化 token 对
type TokenSet struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
	Scope        string    `json:"scope"`
	ClientID     string    `json:"client_id"`
	ServerURL    string    `json:"server_url"`
}

// IsExpired token 是否已过期 (提前 30 秒视为过期)
func (t *TokenSet) IsExpired() bool {
	return time.Now().After(t.ExpiresAt.Add(-30 * time.Second))
}

// ClientRegistration 动态注册响应
type ClientRegistration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret,omitempty"`
}

// ---------- Managed Models ----------

// ManagedModel 托管模型
type ManagedModel struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Provider  string `json:"provider"`
	ModelID   string `json:"modelId"`
	MaxTokens int    `json:"maxTokens"`
	IsEnabled bool   `json:"isEnabled"`
}

// ChatMessage 聊天消息
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatRequest 聊天请求
type ChatRequest struct {
	Messages  []ChatMessage `json:"messages"`
	Stream    bool          `json:"stream,omitempty"`
	MaxTokens int           `json:"max_tokens,omitempty"`
}

// ChatResponse 同步聊天响应
type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Choices []struct {
		Index   int         `json:"index"`
		Message ChatMessage `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

// StreamEvent SSE 流式事件
type StreamEvent struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

// ---------- Entitlements ----------

// EntitlementBalance 权益余额
type EntitlementBalance struct {
	TotalTokenQuota     int64 `json:"totalTokenQuota"`
	TotalTokenUsed      int64 `json:"totalTokenUsed"`
	TotalTokenRemaining int64 `json:"totalTokenRemaining"`
	TotalCallQuota      int   `json:"totalCallQuota"`
	TotalCallUsed       int   `json:"totalCallUsed"`
	TotalCallRemaining  int   `json:"totalCallRemaining"`
	ActiveEntitlements  int   `json:"activeEntitlements"`
}

// ---------- API Response Wrapper ----------

// APIResponse nexus-v4 标准响应
type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
