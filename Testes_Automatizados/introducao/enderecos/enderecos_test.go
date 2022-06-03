package enderecos

import "testing"

type cecnarioDeTeste struct {
	enderecoInserido string
	enderecoEsperado string
}

func TestTipoDeEnderecos(t *testing.T) {

	cenariosDeTeste := []cecnarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia Paulista", "Rodovia"},
		{"Praça Paulista", "Tipo Inválido"},
		{"Estrada Paulista", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA QUALQUER", "Avenida"},
		{"", "Tipo Inválido"},
		{"123", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecosRecebidos := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecosRecebidos != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				cenario.enderecoEsperado,
				cenario.enderecoInserido,
			)
		}

	}

}

//Antes de Refatorar
// package enderecos

// import "testing"

// func TestTipoDeEnderecos(t *testing.T) {
// 	//enderecoParaTeste := "Avenida Paulista"
// 	enderecoParaTeste := "Rua ABC"
// 	tipoDeEnderecoEsperado := "Avenida"
// 	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

// 	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
// 		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
// 			tipoDeEnderecoEsperado,
// 			tipoDeEnderecoRecebido,
// 		)
// 	}

// }
