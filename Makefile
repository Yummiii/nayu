build:
	./tailwindcss -i web/assets/css/input.css -o dist/tailwind.css --minify
	templ generate
	go build -o dist/nayu main.go

run: build
	go run main.go
