module github.com/denysvitali/mbstatus

go 1.19

require github.com/denysvitali/go-mobilebroadband v1.0.0

require github.com/godbus/dbus/v5 v5.0.6 // indirect

replace (
	github.com/denysvitali/go-mobilebroadband v1.0.0 => ../go-mobilebroadband
)
