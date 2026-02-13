# Pipe Me

A CLI tool for converting directory structures into a readable markdown structure to your terminal or clipboard

```
├── folder1
│   ├── file1.txt
│   └── file2.txt
└── folder2
    └── file3.csv
```

## Install
```bash
go install github.com/ColinEge/pipeme@latest
```

## Example usage
Show directories in current directory (recursive)
```bash
pipeme
```
### flags
- `-d <dir>` directory to start search
- `-i dir1,dir2` paths to ignore
- `-f` show files as well as directories
- `-c` copy tree to clipboard

### full example
```bash
pipeme -d ./path -i .git,.gitignore,node_modules -f -c
```