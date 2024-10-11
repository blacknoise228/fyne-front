androidbuild:
	fyne package -os android -appID com.example.myapp

build: 
	go build -o myapp

.PHONY: androidbuild build