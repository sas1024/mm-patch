package patch

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"os"
	"strings"
	"time"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAsSxefJF8G0o6BJI5/cSHo74MQGgThowoVVm5orMfRs8KPRWm
LgZRM2WKl/KRVUT51VXc1NEma2gobleBION5JECFfKjo4GUE4T8bk2bIiYWPP0+q
UQewl2g7j5IENWMtVBPLYYx8c68SVgnkjkSIG7pr/YGN17J5z7KXc71XjYuWDBel
8f0YAz2KBn58ZcpQnu9K6ajIs7YjJkuni3y8dmvQro8LSuLRtWRCjoFy93VBJAF8
UPwCKBHJVkvw9cdwOpDKtfMWGGKf0Ks8hZ7C2Q1pE6btSfc67mxHR9Sva/aqbBzf
w4TD80Z5JMbJYsclqM7aKEwlZJMgJasPSCtGmQIDAQABAoIBACX55DDgQFFbKi9z
pnGCEC7lXBvsEw9aeIS+7D73FQOo+kFYpBumZ/5TzA7AxC0aUVDMjD6jrBAGre/k
2r1RdNRz7gjn7a63iIG9dKw2MlLj6W4BJfkjZFM32NhvzG4jGYK1kXkR01U6l/wQ
N8jU4LXM0jvu6pfq3hhKaBM7aQpiHiZXmRVyiog5jXG8HKmbzsghCM/uyoIcC792
X9ViRAdG/FZWl5Q1+uwijy7wPadpnTNMy9QVqsEG3pdZriXsBFHvVHKA8g01B2Io
eEsB6kUGZs2stoWeNfxIOuLcD1S/N9VP2luo+59J6jzwS5gS1ZKDXMkyfg5/eS9s
gREs03kCgYEA+iFEhJ9bWHYvusLih2C5ecx26K66d75tve/4rGYE/2j0DDfbZjNi
eL8ijgXfDB8Gg3SEUPbUSz++zGmY0iIulmVmP+evdh/dhbC5fLTZTqWwzUJMRbMl
UWMhDNgo9ypYtYGLTfge0SCPNWVnIrNvc1Heleh1wCgPZDFSwVAOz9UCgYEAtVTK
xp8m0cUglhTYf6wUP79PqdwZ31WDVYDI4lkacWLiB2A5Tzl14nMuhM4v7T5olp10
216lvE6veyz4arDSqWmm0ncOQ6xTDGrBjEmzgef1EqLYslElVIlispNV/bqyEwyG
g30Q1n03xRXKakZwJoKx/XnPyzVTILUDdE8YgbUCgYBkrrdx2uNd/FTCDGg6rgh6
qn7CsnKEeLab2dhzLK2eUZTKxkEeJljg2a8DFAHwUxzAFUqdfH1/vK0EgwzsqK2w
BCjgWFYcaj807Sn8tJ80NSWxuZoSBEZlOE25adkzhGwow3hbbiCZdU2v5J1bLncS
KEY8eVHMg1OOtPvmrF8J8QKBgAhvpnNxKhQuUKLK23utHNAObX1gkQ+T4eVTdYUa
UiGeURe4wVHPQY3EgBCLqy0lbyY6sxoVoC5PlthrMi98hIB/OtSl11MMrFxyhwio
0SIlEYDJdL1vCwaQ0bevJRwF2I0MUyHA6syfzL1tkxo4prUT9YXuad1xYKmv4jZC
C8jVAoGAQ2G7grhSKMFOvaajhpmj7dZBJ9DpB8uehuKT+Taqpl8v7VnmvoDxWFuF
GZF9QAU6acl1llsuuHWV5thBWNQ3Y6OLm6s/yBF2uQL1viRI4/bWHqE2BM+LWbWf
SRFdcEOBgM/1K64u7TtzxU4brWP2VNlX0rzoyvLZLpy+I+BVTVw=
-----END RSA PRIVATE KEY-----`

type Customer struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Company string `json:"company"`
}

type Features struct {
	Users                     int  `json:"users"`
	Ldap                      bool `json:"ldap"`
	LdapGroups                bool `json:"ldap_groups"`
	Mfa                       bool `json:"mfa"`
	GoogleOauth               bool `json:"google_oauth"`
	Office365Oauth            bool `json:"office365_oauth"`
	Compliance                bool `json:"compliance"`
	Cluster                   bool `json:"cluster"`
	Metrics                   bool `json:"metrics"`
	Mhpns                     bool `json:"mhpns"`
	Saml                      bool `json:"saml"`
	ElasticSearch             bool `json:"elastic_search"`
	Announcement              bool `json:"announcement"`
	ThemeManagement           bool `json:"theme_management"`
	EmailNotificationContents bool `json:"email_notification_contents"`
	DataRetention             bool `json:"data_retention"`
	MessageExport             bool `json:"message_export"`
	CustomPermissionsSchemes  bool `json:"custom_permissions_schemes"`
	CustomTermsOfService      bool `json:"custom_terms_of_service"`
	GuestAccounts             bool `json:"guest_accounts"`
	GuestAccountsPermissions  bool `json:"guest_accounts_permissions"`
	IdLoaded                  bool `json:"id_loaded"`
	LockTeammateNameDisplay   bool `json:"lock_teammate_name_display"`
	Cloud                     bool `json:"cloud"`
	SharedChannels            bool `json:"shared_channels"`
	RemoteClusterService      bool `json:"remote_cluster_service"`
	Openid                    bool `json:"openid"`
	EnterprisePlugins         bool `json:"enterprise_plugins"`
	AdvancedLogging           bool `json:"advanced_logging"`
	FutureFeatures            bool `json:"future_features"`
}

type License struct {
	ID           string   `json:"id"`
	IssuedAt     int64    `json:"issued_at"`
	StartsAt     int64    `json:"starts_at"`
	ExpiresAt    int64    `json:"expires_at"`
	SkuName      string   `json:"sku_name"`
	SkuShortName string   `json:"sku_short_name"`
	Customer     Customer `json:"customer"`
	Features     Features `json:"features"`
	IsTrial      bool     `json:"is_trial"`
	IsGovSku     bool     `json:"is_gov_sku"`
}

type LicenseCustomize struct {
	Name       string
	Email      string
	Company    string
	Users      int
	ExpireYear int
}

func defaultLicense(licenseCustomize LicenseCustomize) License {
	return License{
		ID:           GenerateRandomLowerString(26),
		IssuedAt:     time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC).UnixMilli(),
		StartsAt:     time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC).UnixMilli(),
		ExpiresAt:    time.Date(licenseCustomize.ExpireYear, 12, 31, 23, 59, 59, 0, time.UTC).UnixMilli(),
		SkuName:      GenerateRandomString(14),
		SkuShortName: "enterprise",
		Customer: Customer{
			ID:      GenerateRandomLowerString(26),
			Name:    licenseCustomize.Name,
			Email:   licenseCustomize.Email,
			Company: licenseCustomize.Company,
		},
		Features: Features{
			Users:                     licenseCustomize.Users,
			Ldap:                      true,
			LdapGroups:                true,
			Mfa:                       true,
			GoogleOauth:               true,
			Office365Oauth:            true,
			Compliance:                true,
			Cluster:                   true,
			Metrics:                   true,
			Mhpns:                     true,
			Saml:                      true,
			ElasticSearch:             true,
			Announcement:              true,
			ThemeManagement:           true,
			EmailNotificationContents: true,
			DataRetention:             true,
			MessageExport:             true,
			CustomPermissionsSchemes:  true,
			CustomTermsOfService:      true,
			GuestAccounts:             true,
			GuestAccountsPermissions:  true,
			IdLoaded:                  true,
			LockTeammateNameDisplay:   true,
			Cloud:                     false,
			SharedChannels:            true,
			RemoteClusterService:      true,
			Openid:                    true,
			EnterprisePlugins:         true,
			AdvancedLogging:           true,
			FutureFeatures:            true,
		},
		IsTrial:  false,
		IsGovSku: false,
	}
}

func GenerateLicense(licenseUpdate LicenseCustomize, fname string) error {
	license := defaultLicense(licenseUpdate)

	b, _ := json.Marshal(license)
	jsonBytes := compactJson(b)

	block, _ := pem.Decode([]byte(privateKey))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	hashed := sha512.Sum512(jsonBytes)

	signature, _ := rsa.SignPKCS1v15(nil, key, crypto.SHA512, hashed[:])

	outBytes := append(jsonBytes, signature...)

	licenseKey := base64.StdEncoding.EncodeToString(outBytes)

	return os.WriteFile(fname, []byte(licenseKey), os.ModePerm)
}

func compactJson(in []byte) []byte {
	out := strings.Replace(string(in), ": ", ":", -1)
	out = strings.Replace(out, ", ", ",", -1)
	return []byte(out)
}
