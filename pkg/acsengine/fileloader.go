package acsengine

//go:generate go-bindata -pkg $GOPACKAGE -prefix ../../parts/ -o templates.go ../../parts/
// fileloader use go-bindata (https://github.com/jteeuwen/go-bindata)
// go-bindata is the way we handle embedded files, like binary, template, etc.
