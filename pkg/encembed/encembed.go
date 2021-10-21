package encembed

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
	"text/template"

	"filippo.io/age"
)

func KeyGen() string {
	//base64 key
	keybytes := make([]byte, 32)
	rand.Read(keybytes)
	return base64.RawStdEncoding.EncodeToString(keybytes)
}

type Config struct {
	PkgName          string
	FuncName         string
	Key              string
	EmbedName        string
	EncryptedVarName string
	DecryptedVarName string
	ExternalKey      string

	Infile  string
	Outfile string
}

func Embed(cfg Config, byts []byte) error {
	var inf io.Reader
	var err error
	if byts == nil {
		inf, err = os.Open(cfg.Infile)
		if err != nil {
			return err
		}
	}

	of, err := os.Create(cfg.EmbedName)
	if err != nil {
		return err
	}

	rcpt, err := age.NewScryptRecipient(cfg.Key)
	if err != nil {
		return err
	}
	w, err := age.Encrypt(of, rcpt)
	if err != nil {
		return err
	}
	io.Copy(w, inf)
	w.Close()

	srcf, err := os.Create(cfg.Outfile)
	if err != nil {
		return err
	}

	tmp, err := template.New("encthing").Parse(tpl)
	if err != nil {
		return err
	}
	err = tmp.Execute(srcf, cfg)
	if err != nil {
		return err
	}
	srcf.Close()
	if cfg.ExternalKey != "" {
		kf, err := os.Create(cfg.ExternalKey)
		if err != nil {
			return err
		}
		kf.WriteString(cfg.Key)
		kf.Close()
	}
	return nil
}
