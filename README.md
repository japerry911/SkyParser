# SkyParser

- Basic CLI app that reads a specified CSV file and seeks either 10 or a specified number of lines and prints them out 
in indented/marshalled format.

## Instructions
- if building (from root directory of repo):
    ```bash
      go build main.go
    ```

- if running/post build:
  ```bash
    ./main -filename=<filename> -lines=<optional-integer>
  ```
