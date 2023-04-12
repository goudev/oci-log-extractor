# Makefile para compilar o programa main.go para várias plataformas

# Configurações para o compilador Go
GO=go
BUILD=go build
DIST=dist

# Nome do arquivo de saída
NAME=Think-OciLogExtractor

# Compilação para Windows 64 bits
windows:
	env GOOS=windows GOARCH=amd64 $(BUILD) -o $(DIST)/$(NAME)-windows-amd64.exe main.go

# Compilação para Linux 64 bits (x86)
linux-x86:
	env GOOS=linux GOARCH=amd64 $(BUILD) -o $(DIST)/$(NAME)-linux-x86_64 main.go

# Compilação para Linux 64 bits (ARM)
linux-arm64:
	env GOOS=linux GOARCH=arm64 $(BUILD) -o $(DIST)/$(NAME)-linux-arm64 main.go

# Compilação para MacOS M1
macos-m1:
	env GOOS=darwin GOARCH=arm64 $(BUILD) -o $(DIST)/$(NAME)-macos-arm64 main.go

# Compilação para MacOS Intel
macos-intel:
	env GOOS=darwin GOARCH=amd64 $(BUILD) -o $(DIST)/$(NAME)-macos-x86_64 main.go

# Compilação para todas as plataformas
all: windows linux-x86 linux-arm64 macos-m1 macos-intel

# Comando para limpar os arquivos gerados
clean:
	rm -rf $(DIST)/*