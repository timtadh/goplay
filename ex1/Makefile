

main:
	6g -o main.6 main.go
build: main
	6l -o main main.6
run: build
	./main

.PHONY : clean
clean :
		-rm main *.6
		ls
