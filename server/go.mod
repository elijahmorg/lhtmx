module github.com/elijahmorg/lhtmxserver

go 1.22.1

require github.com/gofiber/fiber/v2 v2.52.2

require github.com/chasefleming/elem-go v0.25.0 // indirect

require github.com/elijahmorg/lhtmx v0.0.0

replace github.com/elijahmorg/lhtmx v0.0.0 => ../lhtmx

require (
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/klauspost/compress v1.17.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.51.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
