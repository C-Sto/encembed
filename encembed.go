package main

import (
	"crypto/rand"
	"encoding/base64"
	"flag"
	"io"
	"os"
	"text/template"

	"filippo.io/age"
)

type encembedtpl struct {
	PkgName   string
	FuncName  string
	Key       string
	EmbedName string
}

func main() {
	cfg := encembedtpl{}
	infname := flag.String("i", "", "input file")
	flag.StringVar(&cfg.EmbedName, "o", "encembedded", "encrypted output file")
	ofsrcname := flag.String("srcname", "zencembed.go", "source file name to create")
	flag.StringVar(&cfg.FuncName, "funcname", "embedded", "name of function to return decrypted input file")
	flag.StringVar(&cfg.PkgName, "pkgname", "main", "name of package for source file to output")
	flag.Parse()

	cfg.Key = keyGen()
	inf, err := os.Open(*infname)
	if err != nil {
		panic(err)
	}
	of, err := os.Create(cfg.EmbedName)
	if err != nil {
		panic(err)
	}

	rcpt, err := age.NewScryptRecipient(cfg.Key)
	if err != nil {
		panic(err)
	}
	w, err := age.Encrypt(of, rcpt)
	if err != nil {
		panic(err)
	}
	io.Copy(w, inf)
	w.Close()

	srcf, err := os.Create(*ofsrcname)
	if err != nil {
		panic(err)
	}

	tmp, err := template.New("encthing").Parse(tpl)
	if err != nil {
		panic(err)
	}
	err = tmp.Execute(srcf, cfg)
	if err != nil {
		panic(err)
	}
	srcf.Close()

}

const tpl = `
package {{.PkgName}}

import(
	"filippo.io/age"
	"io"
	_ "embed"
	"bytes"
)

//go:embed {{.EmbedName}}
var emb []byte
func {{.FuncName}}() []byte {
	i, _ := age.NewScryptIdentity("{{.Key}}")
	r, _ := age.Decrypt(bytes.NewReader(emb), i)
	a, _ := io.ReadAll(r)
	return a
}
`

func keyGen() string {
	//base64 key
	keybytes := make([]byte, 32)
	rand.Read(keybytes)
	return base64.RawStdEncoding.EncodeToString(keybytes)
}
