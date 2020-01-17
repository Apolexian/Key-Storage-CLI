Key Storage Secret CLI
------------------------

A basic way to store API keys on local machine. Provides limited security by
storing using encryption and not plain text (for more info see
```internal/encryption/encryption.go```). The motivation for this
project was my personal need for an offline storage method for API keys
as access to web storage (e.g https://www.vaultproject.io/) was not always available.
Basic encryption was developed as the idea of storing keys as plaintext was unsatisfactory.

Project Dependencies
---------------------
The project was built with and tested on: <br>

|   Go	| 1.10.4|
|---	|---	|
|   **OS**	|**debian/amd64**|   

Project dependencies outside of stdlib: <br>
Cobra CLI tool - https://github.com/spf13/cobra <br>
Home DIR location - https://github.com/mitchellh/go-homedir <br>


Usage
-------------
Clone repo or download zip to desired directory. E.g YourDirectory. <br>
Run the make file. <br>
```bash
[~/YourDir]: make
```
api_storage will be created in the cmd directory. Change into cmd
directory and use from there. <br>
To set a key for an API:
```bash
[~/YourDir/cmd]: ./api_storage set [API name] [API key]
```
Will set a name:key pair in the .secrets folder in the home directory. <br>
To retrieve a key: 
```bash
[~/YourDir/cmd]: ./api_storage get [API name]
```
Will retrieved the corresponding key. 

To generate HTML documentation for source code: <br>
Change directory to needed one and run the Makefile. For example
if documentation is needed for encryption and storage:
```bash
[~/YourDir/internal]: make
```
Road-map
-------
Notes on code style:<br>
* gofmt was used to format source code, read more
at https://golang.org/cmd/gofmt/
* logger format : ```log.Ldate|log.Ltime|log.Lshortfile```

Linting:<br>
* golint was used, see more at https://github.com/golang/lint

Future goals:<br>
* include more functionality (e.g delete a key)
* improve deployment so script can be usable from anywhere
* add possibility of specifying .secret directory


