# OCILogExtractor

OCILogExtractor is a command-line tool (CLI) that extracts logs from Oracle Cloud Infrastructure (OCI) from a start date to an end date. This tool is useful for analyzing logs for issues or for monitoring activities in your OCI account.

## Requirements

* Go 1.16 or higher installed on the machine.
* An OCI account with access to the logs service configured and with CLI installed. To learn how to configure and install the CLI, please visit: https://docs.oracle.com/pt-br/iaas/Content/API/SDKDocs/cliinstall.htm

## Installation

1. Clone the repository to your local machine:
`https://github.com/goudev/oci-log-extractor.git`


2. Navigate to the project folder:
`cd oci-log-extractor`

3. Compile the executables with the following command:
`make all`

## How to execute?
To use OCILogExtractor, run the following command:

`./OCILogExtractor-<os>-<arch> -data-inicial=(formato 2023-01-31T00:00:00) -data-final=(formato 2023-01-31T00:00:00) -compartment=(ocid1.compartment.oc1..xxxxxxxxxxxxxxx)`

<br>
Replace `(2023-01-31T00:00:00)` with the desired start date and time, and `(2023-01-31T00:00:00)` with the desired end date and time. <br>
Replace `(ocid1.compartment.oc1..xxxxxxxxxxxxxxx)` with the ID of the compartment from which you want to extract the logs.

<br>
The extracted logs will be saved to a file named *log-2023-01-31T00-00-00_2023-01-31T00-00-00.jsonl* in the current directory.

<br>

## Downloadable executables

* [Linux - arm64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-linux-arm64)
* [Linux - x86_64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-linux-x86_64)
* [MacOS - arm64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-macos-arm64)
* [MacOS - x86_64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-macos-x86_64)
* [Windows - amd64](https://objectstorage.sa-saopaulo-1.oraclecloud.com/n/gr1ezkxpdb2l/b/bucket-20230411-2326/o/OciLogExtractor-windows-amd64.exe)
