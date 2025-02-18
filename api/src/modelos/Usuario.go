package modelos

import  (
	"errors"
        "time"
        "strings"
        )

type Usuario struct {
  ID uint64             `json:"id,omitempty"`
  Nome string           `json:"nome,omitempty"`
  Nick string           `json:"nick,omitempty"`
  Email string          `json:"email,omitempty"`
  Senha string          `json:"senha,omitempty"`
  CriadoEm time.Time    `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) Preparar() error {
  if erro := usuario.validar(); erro != nil {
    return erro
  }
  
  usuario.formatar()
  return nil
}

func (usuario *Usuario) validar() error {
  if usuario.Nome == "" {
    return errors.New("nome de usuario em branco")
  }
  
  if usuario.Nick == "" {
    return errors.New("nick de usuario em branco")
  }
  
  if usuario.Email == "" {
    return errors.New("email de usuario em branco")
  }
  
  if usuario.Senha == "" {
    return errors.New("senha de usuario em branco")
  }
  
  return nil
}

func (usuario *Usuario) formatar() {
  usuario.Nome = strings.TrimSpace(usuario.Nome)
  usuario.Nick = strings.TrimSpace(usuario.Nick)
  usuario.Email = strings.TrimSpace(usuario.Email)
}
