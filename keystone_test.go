package keystone

import (
	"fmt"
	"testing"
)

const (
	badCert = `-----BEGIN CERTIFICATE-----
MIIC+DCCAeACCQD6yTQ6qQbuNjANBgkqhkiG9w0BAQUFADA+MQswCQYDVQQGEwJV
UzELMAkGA1UECBMCQ0ExEDAOBgNVBAoTB0NvbXBhbnkxEDAOBgNVBAMTB0NvbXBh
bnkwHhcNMTYwMjExMDcwODUwWhcNMjYwMjA4MDcwODUwWjA+MQswCQYDVQQGEwJV
UzELMAkGA1UECBMCQ0ExEDAOBgNVBAoTB0NvbXBhbnkxEDAOBgNVBAMTB0NvbXBh
bnkwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQDemgoeEl7B+B7v+mE3
zYYHckHJyHWniQeubebjQGOJazB5FF0jTbPW4ipyleozNzOjHah54b7gPFTriu4P
50of3tNfVG5/E1NNAFJ+3cRK/xDM/X4b8ofYsQ0eQVJHEm5cc0aOW/CVaMAWSzwm
sAr5gZ6nfa6EO1Dm42ODlxvVRiwi6+MW/3QQkPdFDSz8WbcqFMH/aj1PD8m3gVGE
mUFL3kui3r+7KR+gT8fhy5Oev3nOYm1lVVFh3S2/Yw7MBFul15dC40+O68kXTwW9
wTJh2QOjlIOplqeqF/4I0m8NU6ik9PF+y1nGQLevLGWmYNKNUdsDhLbi5cExgB/Z
HCC3AgMBAAEwDQYJKoZIhvcNAQEFBQADggEBAFR2K3Z0UDKe9KI6hamjHDjS7fAk
8f9LJTjPDqo75iW4mvJUMI/kqDx7+5N/0fw9qwmWu6giC4VvNQl1YfWaRNQKw6zK
bbfjEbSfW7XXA6r/f8DyjlEVMIK+JIILcP6yB/7hoDV0RXmJfp/BTYKQqvdS2z7y
bDp3F0KbyRW9LJ7F+6pHwEZuSSuZxA6K/ZQ5eWzP2/lmAlmH9mJ4ZcOEw5b7btWb
U0sWyvLO550nKzwvAJI7LhX40AWa3pppOO/eXyw01pzK5jV9lNk/5MukHfzaSYpm
NyXHkZoGalzc5Xlr3HpG9GYjttoS7+hFgcP5JZkWX/9yU=
-----END CERTIFICATE-----`
	goodCert = `-----BEGIN CERTIFICATE-----
MIIG9zCCBd+gAwIBAgIIQgaurK/tlvEwDQYJKoZIhvcNAQELBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTYwNTA0MDkwNTU2WhcNMTYwNzI3MDgzOTAw
WjBmMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEVMBMGA1UEAwwMKi5n
b29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE76rXsthJIywlimmp
ZEp+ah6GNUK0b2GHpDWvenIjlwcVPrdAAD36xA8EX2qXP4U/HJilnxm5OsrdZg0E
yNHAl6OCBI8wggSLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjCCA04G
A1UdEQSCA0UwggNBggwqLmdvb2dsZS5jb22CDSouYW5kcm9pZC5jb22CFiouYXBw
ZW5naW5lLmdvb2dsZS5jb22CEiouY2xvdWQuZ29vZ2xlLmNvbYIWKi5nb29nbGUt
YW5hbHl0aWNzLmNvbYILKi5nb29nbGUuY2GCCyouZ29vZ2xlLmNsgg4qLmdvb2ds
ZS5jby5pboIOKi5nb29nbGUuY28uanCCDiouZ29vZ2xlLmNvLnVrgg8qLmdvb2ds
ZS5jb20uYXKCDyouZ29vZ2xlLmNvbS5hdYIPKi5nb29nbGUuY29tLmJygg8qLmdv
b2dsZS5jb20uY2+CDyouZ29vZ2xlLmNvbS5teIIPKi5nb29nbGUuY29tLnRygg8q
Lmdvb2dsZS5jb20udm6CCyouZ29vZ2xlLmRlggsqLmdvb2dsZS5lc4ILKi5nb29n
bGUuZnKCCyouZ29vZ2xlLmh1ggsqLmdvb2dsZS5pdIILKi5nb29nbGUubmyCCyou
Z29vZ2xlLnBsggsqLmdvb2dsZS5wdIISKi5nb29nbGVhZGFwaXMuY29tgg8qLmdv
b2dsZWFwaXMuY26CFCouZ29vZ2xlY29tbWVyY2UuY29tghEqLmdvb2dsZXZpZGVv
LmNvbYIMKi5nc3RhdGljLmNugg0qLmdzdGF0aWMuY29tggoqLmd2dDEuY29tggoq
Lmd2dDIuY29tghQqLm1ldHJpYy5nc3RhdGljLmNvbYIMKi51cmNoaW4uY29tghAq
LnVybC5nb29nbGUuY29tghYqLnlvdXR1YmUtbm9jb29raWUuY29tgg0qLnlvdXR1
YmUuY29tghYqLnlvdXR1YmVlZHVjYXRpb24uY29tggsqLnl0aW1nLmNvbYIaYW5k
cm9pZC5jbGllbnRzLmdvb2dsZS5jb22CC2FuZHJvaWQuY29tggRnLmNvggZnb28u
Z2yCFGdvb2dsZS1hbmFseXRpY3MuY29tggpnb29nbGUuY29tghJnb29nbGVjb21t
ZXJjZS5jb22CCnVyY2hpbi5jb22CCnd3dy5nb28uZ2yCCHlvdXR1LmJlggt5b3V0
dWJlLmNvbYIUeW91dHViZWVkdWNhdGlvbi5jb20wCwYDVR0PBAQDAgeAMGgGCCsG
AQUFBwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJ
QUcyLmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20v
b2NzcDAdBgNVHQ4EFgQUu7ZNID1gvqDryLWKCZeHA5p3NywwDAYDVR0TAQH/BAIw
ADAfBgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAhBgNVHSAEGjAYMAwG
CisGAQQB1nkCBQEwCAYGZ4EMAQICMDAGA1UdHwQpMCcwJaAjoCGGH2h0dHA6Ly9w
a2kuZ29vZ2xlLmNvbS9HSUFHMi5jcmwwDQYJKoZIhvcNAQELBQADggEBAJYLxnTV
9VJWrpPWUIcAcFbOffFjvfgmyW14PX5mI70mw1UhZQiAMSm/nkihmgbkr4VTlyiX
R6gINdsluNwdleFptwX6Lu9tcwnkJC740OsDMiCQCh9pQkdS0e4/CqfpukWDx9fZ
6S5UWXfFUURJoTeXHo6DRWiU/etT5NnPVgSZP4NDGF9o/86lBCB216II3+d8mOTs
cTHzc9JraiEs8Wlyk6RrmqBNTS/unWLzOFbbIJsQZYENyUyviJtxqsHDfCV2RAA0
sxhac+mgQNMDdqVqvx705n7l5TMGK/JNNr0q+XnOl0Gxl9mS8AQ3MORj34/pXyDC
dG7l3S70kqtlvDM=
-----END CERTIFICATE-----`
	badCert2 = `-----BEGIN CERTIFICATE-----
MIIHCDCCBfCgAwIBAgIQTNer/7wzBW3WI/MKERuV+TANBgkqhkiG9w0BAQsFADB3
MQswCQYDVQQGEwJVUzEdMBsGA1UEChZUU3ltYW50ZWMgQ29ycG9yYXRpb24xHzAd
BgNVBAsTFlN5bWFudGVjIFRydXN0IE5ldHdvcmsxKDAmBgNVBAMTH1N5bWFudGVj
IENsYXNzIDMgRVYgU1NMIENBIC0gRzMwHhcNMTYwMzI4MDAwMDAwWhcNMTcxMDE1
MjM1OTU5WjCCARgxEzARBgsrBgEEAYI3PAIBAxMCVVMxGzAZBgsrBgEEAYI3PAIB
AgwKQ2FsaWZvcm5pYTEdMBsGA1UEDxMUUHJpdmF0ZSBPcmdhbml6YXRpb24xETAP
BgNVBAUTCEMwODA2NTkyMQswCdYDVQQGEwJVUzEOMAwGA1UEEQwFOTUwMTQxEzAR
BgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAcMCUN1cGVydGlubzEYMBYGA1UECQwP
MSBJbmZpbml0ZSBMb2awMRMwEQYDVQQKDApBcHBsZSBJbmMuMSUwIwYDVQQLDBxJ
bnRlcm5ldCBTZXJ2aWNlcyBmb3IgQWthbWFpMRYwFAYDVQQDDA13d3cuYXBwbGUu
Y29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6I2D/ncBDY/lKFFg
wgJjhg5n/80vlupJdS08BEBI0yJvgkJGh3rYCsPIpMBmAR8vE+08UngxF+9ShQFC
d8xzHIRdEFHwy8zEbC5l+0Ay4vYpYzqM/oD5kODoRXNWs+vBxE6H0qH6IkUq/sXE
rIAJNROc3h2yGeWXUe2D9ZnrFOZqMcdGC+ismq5erVQi1GJivZCP3aH03Uljj0eF
YZLewOzv3HaDmy0UZYGesadC0UBnR8h13hzXpL5mMIBt8uaEhxkQEpcVpjX68vUy
jvgqKyoc8qBWMWayv8gaFbwWR1q6nAsjR79yQMQSgrvnH2A6kj2cue/N/aVPrgxn
TiU/LQIDAQABo4IC6zCCAucwGAYDVR0RBBEwD4INd3d3LmFwcGxlLmNvbTAJBgNV
HRMEAjAAMA4GA1UdDwEB/wQEAwIFoDAdBgNVHSUEFjAUBggrBgEFBQcDAQYIKwYB
BQUHAwIwZgYDVR0gBF8wXTBbBgtghkgBhvhFAQcXBjBMMCMGCCsGAQUFBwIBFhdo
dHRwczovL2Quc3ltY2IuY29tL2NwczAlBggrBgEFBQcCAjAZGhdodHRwczovL2Qu
c3ltY2IuY29tL3JwYTAfBgNVHSMEGDAWgBQBWavn3ToLWaZkY9bPIAdX1ZHnajAr
BgNVHR8EJDAiMCCgHqAchhpodHRwOi8vc3Iuc3ltY2IuY29tL3NyLmNybDBXBggr
BgEFBQcBAQRLMEkwHwYIKwYBBQUHMAGGE2h0dHA6Ly9zci5zeW1jZC5jb20wJgYI
KwYBBQUHMAKGGmh0dHA6wy9zCi5zeW1jYi5jb20vc3IuY3J0MIIBgAYKKwYBBAHW
eQIEAgSCAXAEggFsAWoAdgDd6x0reg1PpipLga2BaHB+Lo6dAdVciI09EcTNtuy+
zAAAAVO91zVxAAAEAwBHMEUCIQDngguDP0qBkp/BEFRSXFa6PyhunIivXR8zTq50
5A8BNgIgS2niWG5EH6Urt9pPu04uIkaMroUDeJNcvZQTZSGkLUIAdwCkuQmQtBhY
FIe7E6LMZ3AKPDWYBPkb37jjd80OyA3cEAAAAVO91zWIAAAEAwBIMEYCIQDJeCdC
wk3B9FfncSyZ9Lf5H06XnV538XHb3psqQ17nOAIhAM2qO3hemAL4jJ25DIb35xDa
ley0Fy/2Iz9/YsIa2PHeAHcAaPaY+B9kga46jO65KB1M/HFRXWeT1ETRCmesu09P
+8QAAAFTvdc1jQAABAMASDBGAiEAiCwdd3kJ4EwgZAVIN9nqXJEI/n11yTfPzNCm
GyehUJoCIQCUAHw7CE81k7yjF0IA4/TU7BDmldzRJzDwx3CfNASKfTANBgkqhkiG
9w0BAQsFAAOCAQEAi72XFUSja4y+H0KbikEojpgkI46bhAU/kkaub36bHGhTPutz
40LUR2QBXaMY3dwePDK2MMK9Frk2x75o1VYFR+VClgWM/BLC2m3/t0RvN7Ftlg0w
QPJvjCZZBdZ6VOZADJJap6EP7I7oJbqeZooNU34h+WydleFjgkHNuLZ+wbR0nzIS
Jy6nSXLd54OHk3yhEPbAUqaxahspaYG9/BJIb6ubXamLaSCxhAnWAk96ieL8tDt+
NkTxkAVdGXe95z6KVnVVqRFZEM19O119gZp7nmLOpYPO9gIwW+iLDxK+w3D3ua5Y
jWLO2LL6V9eoixFRXGkfU83taAwgfFOcEDLC9g==
-----END CERTIFICATE-----`
)

type GetCertTestSet struct {
	input    []byte
	expected string
}

var (
	getCertTests = []GetCertTestSet{
		{input: []byte(goodCert), expected: ""},
		{input: []byte(badCert), expected: "error parsing certificate"},
		{input: []byte(badCert2), expected: "asn1: structure error: tags don't match (6 vs {class:3 tag:22 length:3 isCompound:false}) {optional:false explicit:false application:false defaultValue:<nil> tag:<nil> stringType:0 timeType:0 set:false omitEmpty:false} ObjectIdentifier @2"},
		{input: []byte(nil), expected: "input empty"},
	}
)

func TestGetCert(t *testing.T) {
	for _, test := range getCertTests {
		_, got := GetCert(test.input)
		if got != nil {
			if got.Error() != test.expected {
				fmt.Printf("got: %s\n", got.Error())
				t.Errorf("Expected: %s, got: %s", test.expected, got)
			}
		}
	}
}
