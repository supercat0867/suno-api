package kernel

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"time"
)

type Client struct {
	SessionID    string
	SessionExpAt time.Time
	client       *http.Client
}

// NenClient 实例化客户端
func NenClient() *Client {
	// 从环境变量中获取cookie
	cookieStr := os.Getenv("SUNO_COOKIE")
	if cookieStr == "" {
		panic("请先在环境变量中设置SUNO_COOKIE！")
	}

	// 创建 cookie jar
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	// 创建 HTTP 客户端
	httpClient := &http.Client{
		Jar: jar,
	}
	// 解析 cookie 字符串并添加到 cookie jar
	url, _ := url.Parse("https://clerk.suno.com")
	cookies := parseCookies(cookieStr)
	jar.SetCookies(url, cookies)

	// 获取sessionId
	var client Client
	client.client = httpClient
	resp, err := client.GetSunoClient()

	if err != nil {
		panic(fmt.Sprintf("sessionId获取失败:%v", err))
	}

	client.SessionExpAt = convertTimestampMillisToTime(resp.Response.Sessions[0].ExpireAt)
	client.SessionID = resp.Response.Sessions[0].Id

	return &client
}

type GetSunoClientResponse struct {
	Response struct {
		Object   string `json:"object"`
		Id       string `json:"id"`
		Sessions []struct {
			Object                   string      `json:"object"`
			Id                       string      `json:"id"`
			Status                   string      `json:"status"`
			ExpireAt                 int64       `json:"expire_at"`
			AbandonAt                int64       `json:"abandon_at"`
			LastActiveAt             int64       `json:"last_active_at"`
			LastActiveOrganizationId interface{} `json:"last_active_organization_id"`
			Actor                    interface{} `json:"actor"`
			User                     struct {
				Id                    string      `json:"id"`
				Object                string      `json:"object"`
				Username              interface{} `json:"username"`
				FirstName             string      `json:"first_name"`
				LastName              string      `json:"last_name"`
				ImageUrl              string      `json:"image_url"`
				HasImage              bool        `json:"has_image"`
				PrimaryEmailAddressId string      `json:"primary_email_address_id"`
				PrimaryPhoneNumberId  interface{} `json:"primary_phone_number_id"`
				PrimaryWeb3WalletId   interface{} `json:"primary_web3_wallet_id"`
				PasswordEnabled       bool        `json:"password_enabled"`
				TwoFactorEnabled      bool        `json:"two_factor_enabled"`
				TotpEnabled           bool        `json:"totp_enabled"`
				BackupCodeEnabled     bool        `json:"backup_code_enabled"`
				EmailAddresses        []struct {
					Id           string `json:"id"`
					Object       string `json:"object"`
					EmailAddress string `json:"email_address"`
					Reserved     bool   `json:"reserved"`
					Verification struct {
						Status   string      `json:"status"`
						Strategy string      `json:"strategy"`
						Attempts interface{} `json:"attempts"`
						ExpireAt interface{} `json:"expire_at"`
					} `json:"verification"`
					LinkedTo []struct {
						Type string `json:"type"`
						Id   string `json:"id"`
					} `json:"linked_to"`
					CreatedAt int64 `json:"created_at"`
					UpdatedAt int64 `json:"updated_at"`
				} `json:"email_addresses"`
				PhoneNumbers     []interface{} `json:"phone_numbers"`
				Web3Wallets      []interface{} `json:"web3_wallets"`
				Passkeys         []interface{} `json:"passkeys"`
				ExternalAccounts []struct {
					Object           string      `json:"object"`
					Id               string      `json:"id"`
					Provider         string      `json:"provider"`
					IdentificationId string      `json:"identification_id"`
					ProviderUserId   string      `json:"provider_user_id"`
					ApprovedScopes   string      `json:"approved_scopes"`
					EmailAddress     string      `json:"email_address"`
					FirstName        string      `json:"first_name"`
					LastName         string      `json:"last_name"`
					AvatarUrl        string      `json:"avatar_url"`
					Username         interface{} `json:"username"`
					PublicMetadata   struct {
					} `json:"public_metadata"`
					Label        interface{} `json:"label"`
					CreatedAt    int64       `json:"created_at"`
					UpdatedAt    int64       `json:"updated_at"`
					Verification struct {
						Status   string      `json:"status"`
						Strategy string      `json:"strategy"`
						Attempts interface{} `json:"attempts"`
						ExpireAt int64       `json:"expire_at"`
					} `json:"verification"`
				} `json:"external_accounts"`
				SamlAccounts   []interface{} `json:"saml_accounts"`
				PublicMetadata struct {
				} `json:"public_metadata"`
				UnsafeMetadata struct {
				} `json:"unsafe_metadata"`
				ExternalId                    interface{}   `json:"external_id"`
				LastSignInAt                  int64         `json:"last_sign_in_at"`
				Banned                        bool          `json:"banned"`
				Locked                        bool          `json:"locked"`
				LockoutExpiresInSeconds       interface{}   `json:"lockout_expires_in_seconds"`
				VerificationAttemptsRemaining int           `json:"verification_attempts_remaining"`
				CreatedAt                     int64         `json:"created_at"`
				UpdatedAt                     int64         `json:"updated_at"`
				DeleteSelfEnabled             bool          `json:"delete_self_enabled"`
				CreateOrganizationEnabled     bool          `json:"create_organization_enabled"`
				LastActiveAt                  int64         `json:"last_active_at"`
				MfaEnabledAt                  interface{}   `json:"mfa_enabled_at"`
				MfaDisabledAt                 interface{}   `json:"mfa_disabled_at"`
				IsTestUser                    bool          `json:"is_test_user"`
				ProfileImageUrl               string        `json:"profile_image_url"`
				OrganizationMemberships       []interface{} `json:"organization_memberships"`
			} `json:"user"`
			PublicUserData struct {
				FirstName       string `json:"first_name"`
				LastName        string `json:"last_name"`
				ImageUrl        string `json:"image_url"`
				HasImage        bool   `json:"has_image"`
				Identifier      string `json:"identifier"`
				ProfileImageUrl string `json:"profile_image_url"`
			} `json:"public_user_data"`
			CreatedAt       int64 `json:"created_at"`
			UpdatedAt       int64 `json:"updated_at"`
			LastActiveToken struct {
				Object string `json:"object"`
				Jwt    string `json:"jwt"`
			} `json:"last_active_token"`
		} `json:"sessions"`
		SignIn              interface{} `json:"sign_in"`
		SignUp              interface{} `json:"sign_up"`
		LastActiveSessionId string      `json:"last_active_session_id"`
		CreatedAt           int64       `json:"created_at"`
		UpdatedAt           int64       `json:"updated_at"`
	} `json:"response"`
	Client interface{} `json:"client"`
}

type GetJwtResponse struct {
	Object string `json:"object"`
	Jwt    string `json:"jwt"`
}

type GenerateSongResponse struct {
	Id    string `json:"id"`
	Clips []struct {
		Id                string      `json:"id"`
		VideoUrl          string      `json:"video_url"`
		AudioUrl          string      `json:"audio_url"`
		ImageUrl          interface{} `json:"image_url"`
		ImageLargeUrl     interface{} `json:"image_large_url"`
		IsVideoPending    bool        `json:"is_video_pending"`
		MajorModelVersion string      `json:"major_model_version"`
		ModelName         string      `json:"model_name"`
		Metadata          struct {
			Tags                     string      `json:"tags"`
			Prompt                   string      `json:"prompt"`
			GptDescriptionPrompt     interface{} `json:"gpt_description_prompt"`
			AudioPromptId            interface{} `json:"audio_prompt_id"`
			History                  interface{} `json:"history"`
			ConcatHistory            interface{} `json:"concat_history"`
			StemFromId               interface{} `json:"stem_from_id"`
			Type                     string      `json:"type"`
			Duration                 interface{} `json:"duration"`
			RefundCredits            interface{} `json:"refund_credits"`
			Stream                   bool        `json:"stream"`
			Infill                   interface{} `json:"infill"`
			HasVocal                 interface{} `json:"has_vocal"`
			IsAudioUploadTosAccepted interface{} `json:"is_audio_upload_tos_accepted"`
			ErrorType                interface{} `json:"error_type"`
			ErrorMessage             interface{} `json:"error_message"`
		} `json:"metadata"`
		IsLiked         bool        `json:"is_liked"`
		UserId          string      `json:"user_id"`
		DisplayName     string      `json:"display_name"`
		Handle          string      `json:"handle"`
		IsHandleUpdated bool        `json:"is_handle_updated"`
		AvatarImageUrl  string      `json:"avatar_image_url"`
		IsTrashed       bool        `json:"is_trashed"`
		Reaction        interface{} `json:"reaction"`
		CreatedAt       time.Time   `json:"created_at"`
		Status          string      `json:"status"`
		Title           string      `json:"title"`
		PlayCount       int         `json:"play_count"`
		UpvoteCount     int         `json:"upvote_count"`
		IsPublic        bool        `json:"is_public"`
	} `json:"clips"`
	Metadata struct {
		Tags                     string      `json:"tags"`
		Prompt                   string      `json:"prompt"`
		GptDescriptionPrompt     interface{} `json:"gpt_description_prompt"`
		AudioPromptId            interface{} `json:"audio_prompt_id"`
		History                  interface{} `json:"history"`
		ConcatHistory            interface{} `json:"concat_history"`
		StemFromId               interface{} `json:"stem_from_id"`
		Type                     string      `json:"type"`
		Duration                 interface{} `json:"duration"`
		RefundCredits            interface{} `json:"refund_credits"`
		Stream                   bool        `json:"stream"`
		Infill                   interface{} `json:"infill"`
		HasVocal                 interface{} `json:"has_vocal"`
		IsAudioUploadTosAccepted interface{} `json:"is_audio_upload_tos_accepted"`
		ErrorType                interface{} `json:"error_type"`
		ErrorMessage             interface{} `json:"error_message"`
	} `json:"metadata"`
	MajorModelVersion string    `json:"major_model_version"`
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
	BatchSize         int       `json:"batch_size"`
}

type GetSongInfoResponse struct {
	Clips []struct {
		Id                string `json:"id"`
		VideoUrl          string `json:"video_url"`
		AudioUrl          string `json:"audio_url"`
		ImageUrl          string `json:"image_url"`
		ImageLargeUrl     string `json:"image_large_url"`
		IsVideoPending    bool   `json:"is_video_pending"`
		MajorModelVersion string `json:"major_model_version"`
		ModelName         string `json:"model_name"`
		Metadata          struct {
			Tags                     string      `json:"tags"`
			Prompt                   string      `json:"prompt"`
			GptDescriptionPrompt     interface{} `json:"gpt_description_prompt"`
			AudioPromptId            interface{} `json:"audio_prompt_id"`
			History                  interface{} `json:"history"`
			ConcatHistory            interface{} `json:"concat_history"`
			StemFromId               interface{} `json:"stem_from_id"`
			Type                     string      `json:"type"`
			Duration                 float64     `json:"duration"`
			RefundCredits            bool        `json:"refund_credits"`
			Stream                   bool        `json:"stream"`
			Infill                   interface{} `json:"infill"`
			HasVocal                 interface{} `json:"has_vocal"`
			IsAudioUploadTosAccepted interface{} `json:"is_audio_upload_tos_accepted"`
			ErrorType                interface{} `json:"error_type"`
			ErrorMessage             interface{} `json:"error_message"`
		} `json:"metadata"`
		IsLiked         bool        `json:"is_liked"`
		UserId          string      `json:"user_id"`
		DisplayName     string      `json:"display_name"`
		Handle          string      `json:"handle"`
		IsHandleUpdated bool        `json:"is_handle_updated"`
		AvatarImageUrl  string      `json:"avatar_image_url"`
		IsTrashed       bool        `json:"is_trashed"`
		Reaction        interface{} `json:"reaction"`
		CreatedAt       time.Time   `json:"created_at"`
		Status          string      `json:"status"`
		Title           string      `json:"title"`
		PlayCount       int         `json:"play_count"`
		UpvoteCount     int         `json:"upvote_count"`
		IsPublic        bool        `json:"is_public"`
	} `json:"clips"`
	NumTotalResults int `json:"num_total_results"`
	CurrentPage     int `json:"current_page"`
}
