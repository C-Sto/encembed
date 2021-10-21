package main

import (
	"flag"

	"github.com/c-sto/encembed/pkg/encembed"
)

func main() {
	cfg := encembed.Config{} //encembedtpl{}
	flag.StringVar(&cfg.Infile, "i", "", "input file")
	flag.StringVar(&cfg.EmbedName, "o", "encembedded", "encrypted output file")
	flag.StringVar(&cfg.Outfile, "srcname", "zencembed.go", "source file name to create")
	flag.StringVar(&cfg.FuncName, "funcname", "embedded", "name of function to return decrypted input file")
	flag.StringVar(&cfg.PkgName, "pkgname", "main", "name of package for source file to output")
	flag.StringVar(&cfg.EncryptedVarName, "encvarname", "cryptembed", "variable name for encrypted resource")
	flag.StringVar(&cfg.DecryptedVarName, "decvarname", "", "variable name to use for decrypted resource (if you don't want to access it via the function)")
	flag.StringVar(&cfg.ExternalKey, "extkey", "", "do not embed the key in the binary (writes to specified filename)")
	flag.StringVar(&cfg.Key, "key", "", "dont generate key, use this one")
	flag.Parse()

	if cfg.ExternalKey != "" && cfg.DecryptedVarName != "" {
		panic("external key and simple var access incompatible")
	}

	if cfg.Key == "" {
		cfg.Key = encembed.KeyGen()
	}

	err := encembed.Embed(cfg, nil)

	if err != nil {
		panic(err)
	}

}
