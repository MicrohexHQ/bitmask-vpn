// Code generated by go generate; DO NOT EDIT.
// This file was generated by vendorize.py
// At 2019-07-30 21:50:50

package config

/* All these constants are defined in the vendor.conf file
 */
const (
	Provider        = "riseup.net"
	ApplicationName = "RiseupVPN"
	BinaryName      = "riseup-vpn"
	DonateURL       = "https://riseup.net/donate"
	HelpURL         = "https://riseup.net/support"
	TosURL          = "https://riseup.net/tos"
	APIURL          = "https://api.black.riseup.net/"
	GeolocationAPI  = "https://api.black.riseup.net:9001/json"
)

/*

CaCert : a string containing a representation of the provider CA, used to
        sign the webapp and openvpn certificates. should be placed in
        config/[provider]-ca.crt

*/
var CaCert = []byte(`-----BEGIN CERTIFICATE-----
MIIFjTCCA3WgAwIBAgIBATANBgkqhkiG9w0BAQ0FADBZMRgwFgYDVQQKDA9SaXNl
dXAgTmV0d29ya3MxGzAZBgNVBAsMEmh0dHBzOi8vcmlzZXVwLm5ldDEgMB4GA1UE
AwwXUmlzZXVwIE5ldHdvcmtzIFJvb3QgQ0EwHhcNMTQwNDI4MDAwMDAwWhcNMjQw
NDI4MDAwMDAwWjBZMRgwFgYDVQQKDA9SaXNldXAgTmV0d29ya3MxGzAZBgNVBAsM
Emh0dHBzOi8vcmlzZXVwLm5ldDEgMB4GA1UEAwwXUmlzZXVwIE5ldHdvcmtzIFJv
b3QgQ0EwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQC76J4ciMJ8Sg0m
TP7DF2DT9zNe0Csk4myoMFC57rfJeqsAlJCv1XMzBmXrw8wq/9z7XHv6n/0sWU7a
7cF2hLR33ktjwODlx7vorU39/lXLndo492ZBhXQtG1INMShyv+nlmzO6GT7ESfNE
LliFitEzwIegpMqxCIHXFuobGSCWF4N0qLHkq/SYUMoOJ96O3hmPSl1kFDRMtWXY
iw1SEKjUvpyDJpVs3NGxeLCaA7bAWhDY5s5Yb2fA1o8ICAqhowurowJpW7n5ZuLK
5VNTlNy6nZpkjt1QycYvNycffyPOFm/Q/RKDlvnorJIrihPkyniV3YY5cGgP+Qkx
HUOT0uLA6LHtzfiyaOqkXwc4b0ZcQD5Vbf6Prd20Ppt6ei0zazkUPwxld3hgyw58
m/4UIjG3PInWTNf293GngK2Bnz8Qx9e/6TueMSAn/3JBLem56E0WtmbLVjvko+LF
PM5xA+m0BmuSJtrD1MUCXMhqYTtiOvgLBlUm5zkNxALzG+cXB28k6XikXt6MRG7q
hzIPG38zwkooM55yy5i1YfcIi5NjMH6A+t4IJxxwb67MSb6UFOwg5kFokdONZcwj
shczHdG9gLKSBIvrKa03Nd3W2dF9hMbRu//STcQxOailDBQCnXXfAATj9pYzdY4k
ha8VCAREGAKTDAex9oXf1yRuktES4QIDAQABo2AwXjAdBgNVHQ4EFgQUC4tdmLVu
f9hwfK4AGliaet5KkcgwDgYDVR0PAQH/BAQDAgIEMAwGA1UdEwQFMAMBAf8wHwYD
VR0jBBgwFoAUC4tdmLVuf9hwfK4AGliaet5KkcgwDQYJKoZIhvcNAQENBQADggIB
AGzL+GRnYu99zFoy0bXJKOGCF5XUXP/3gIXPRDqQf5g7Cu/jYMID9dB3No4Zmf7v
qHjiSXiS8jx1j/6/Luk6PpFbT7QYm4QLs1f4BlfZOti2KE8r7KRDPIecUsUXW6P/
3GJAVYH/+7OjA39za9AieM7+H5BELGccGrM5wfl7JeEz8in+V2ZWDzHQO4hMkiTQ
4ZckuaL201F68YpiItBNnJ9N5nHr1MRiGyApHmLXY/wvlrOpclh95qn+lG6/2jk7
3AmihLOKYMlPwPakJg4PYczm3icFLgTpjV5sq2md9bRyAg3oPGfAuWHmKj2Ikqch
Td5CHKGxEEWbGUWEMP0s1A/JHWiCbDigc4Cfxhy56CWG4q0tYtnc2GMw8OAUO6Wf
Xu5pYKNkzKSEtT/MrNJt44tTZWbKV/Pi/N2Fx36my7TgTUj7g3xcE9eF4JV2H/sg
tsK3pwE0FEqGnT4qMFbixQmc8bGyuakr23wjMvfO7eZUxBuWYR2SkcP26sozF9PF
tGhbZHQVGZUTVPyvwahMUEhbPGVerOW0IYpxkm0x/eaWdTc4vPpf/rIlgbAjarnJ
UN9SaWRlWKSdP4haujnzCoJbM7dU9bjvlGZNyXEekgeT0W2qFeGGp+yyUWw8tNsp
0BuC1b7uW/bBn/xKm319wXVDvBgZgcktMolak39V7DVO
-----END CERTIFICATE-----`)
