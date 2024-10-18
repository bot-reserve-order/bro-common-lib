package bro_domains

type IssueAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type VerifyAccessTokenResponse struct {
	Scope            string `json:"scope"`
	ClientID         string `json:"client_id"`
	ExpiresIn        int    `json:"expires_in"`
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type GetFriendStatusResponse struct {
	FriendFlag bool `json:"friendFlag"`
}
