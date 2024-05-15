build: build-gradient build-laplacian build-sobel

build-gradient:
	@cd gradient \
	&& g++ main.cpp -o main `pkg-config --cflags --libs opencv4` && \
	./main


build-laplacian:
	@cd laplacian \
	&& g++ main.cpp -o main `pkg-config --cflags --libs opencv4` && \
	./main

build-sobel:
	@cd sobel \
	&& g++ main.cpp -o main `pkg-config --cflags --libs opencv4` && \
	./main


