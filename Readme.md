# encembed

Encrypt embedded resource in compiled binary using [age](https://github.com/FiloSottile/age). Meant for usage with `go generate`.

This tool will generate a go source file that embeds an encrypted version of the file that is specified, and provides a function to access the plaintext content of that file. Options allow for arbitrary naming of the function, encrypted file, and optionally the ability to not include the password in the file.

```
  -decvarname string
        variable name to use for decrypted resource (if you don't want to access it via the function)
  -encvarname string
        variable name for encrypted resource (default "cryptembed")
  -extkey string
        do not embed the key in the binary (writes to specified filename)
  -funcname string
        name of function to return decrypted input file (default "embedded")
  -i string
        input file
  -o string
        encrypted output file (default "encembedded")
  -pkgname string
        name of package for source file to output (default "main")
  -srcname string
        source file name to create (default "zencembed.go")
```

Examples can be found in the example dir, but basic usage to replace your original embed is:

This will encrypt `plaintext_file.txt`, and make a variable called `plaintext` available that exposes the decrypted data.

```
func main(){
    log.Println(string(plaintext))
}
//go:generate go run github.com/c-sto/encembed -i plaintext_file.txt -decvarname plaintext
```

todo: better readme
