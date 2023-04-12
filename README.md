# OCILogExtractor

OCILogExtractor é uma ferramenta em linha de comando (CLI) que extrai logs do Oracle Cloud Infrastructure (OCI) a partir de uma data inicial até uma data final. Essa ferramenta é útil para analisar logs em busca de problemas ou para monitoramento de atividades em sua conta da OCI.

## Requisitos

- Go 1.16 ou superior instalado na máquina.
- Conta da OCI com acesso ao serviço de logs configurada e com cli instalado. Para saber como configurar e instalar o cli acesse: 
https://docs.oracle.com/pt-br/iaas/Content/API/SDKDocs/cliinstall.htm

## Instalação

1. Clone o repositório em sua máquina local:
`https://github.com/goudev/oci-log-extractor.git`


2. Entre na pasta do projeto:
`cd oci-log-extractor`

3. Compile os executáveis com o seguinte comando:
`make all`

## Como executar? 
Para usar o OCILogExtractor, execute o seguinte comando:

`./OCILogExtractor-<os>-<arch> --data-inicial=(formato 2023-01-31T00:00:00) --data-final=(formato 2023-01-31T00:00:00) --compartment=(ocid1.compartment.oc1..xxxxxxxxxxxxxxx)`

<br>
Substitua `(2023-01-31T00:00:00)` pela data e hora inicial desejada e `(2023-01-31T00:00:00)` pela data e hora final desejada. Substitua `(ocid1.compartment.oc1..xxxxxxxxxxxxxxx)` pelo ID do compartment que você deseja extrair os logs.

<br>
Os logs extraídos serão salvos em um arquivo chamado `log-2023-01-31T00-00-00_2023-01-31T00-00-00.jsonl` no diretório atual.

<br>

## Executáveis para download

* [Linux - arm64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-linux-arm64)
* [Linux - x86_64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-linux-x86_64)
* [MacOS - arm64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-macos-arm64)
* [MacOS - x86_64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-macos-x86_64)
* [Windows - amd64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-windows-amd64.exe)
