# goshasum

CLI tool to create recursively shasum file for directory.

## Usage:

```bash
Usage: ./goshasum [OPTIONS] <path>
OPTIONS:
  -algo string
        Hashing algorithm to use. Possible values: sha1, sha256, sha512 (default "sha256")
  -ignore string
        List of ignored files and directories seperated by common. Example: -ignore=.git,node_modules (default ".git")
  -output string
        Output file name. If missing the output will be printed to STDOUT
```