# goshasum

The goshasum is a CLI tool to create recursively shasum file for directory contents. 

## What is a shasum file?
The "shasum file" is an informal term for a simple text file containing checksums and their corresponding filenames. It can be used to verify the integrity of the files.

Here is the example of the shasum file, called `SHA256FILE`.
```
f5abd3525f5d81bfd987c1cc1c14437f2412201461904ffb1bab4e61e7ebdfb3 .gitignore
c59aee6abac9386dc50b169cdeee09af935d1b84c164688759887aca552f63ac README.md
ef7cd13aee9ab511051b3e927f960bf29ac7a2cf089945b6a2bfbb558e4732db go.mod
76132d0433d5f4c2ee9586311d8544d0356aad7628b9356eb35bd0e7f49c64b0 main.go
```

You can call the Linux command `shasum -c SHA256FILE` to verify the checksums of all the files inside the repository directory.

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