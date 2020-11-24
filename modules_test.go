package hello

// Utiliza-se a dependencia testing, para testar o módulos criado
// Roda-se o comando go test com um arquivo que contenha test no nome.
import "testing"

func TestHello(t *testing.T) {
	// Seta o que é esperado de retorno do modulo
	want := "Olá Mundo."
	// Valida se o retorno é o que esperamos.
	if got := Hello(); got != want {
		// Se for diferente, retorna o que recebemos e o que esperavamos.
		t.Errorf("Hello() = %q, want %q", got, want)
	}

	// Seta o que é esperado de retorno do modulo
	want := "Don't communicate by sharing memory, share memory by communicating."
	// Valida se o retorno é o que esperamos.
	if got := Go(); got != want {
		// Se for diferente, retorna o que recebemos e o que esperavamos.
		t.Errorf("Go() = %q, want %q", got, want)
	}
}
