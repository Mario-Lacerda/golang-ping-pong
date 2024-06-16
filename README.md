# Desafio Dio - Ping Pong com Go



**Códigos para o jogo ping pong na linguagem Go:**

go

```go
// Pacote principal
package main

// Importando as bibliotecas necessárias
import (
    "fmt"
    "math/rand"
    "time"
)

// Constantes para as dimensões da tela
const (
    LarguraTela = 800
    AlturaTela = 600
)

// Estrutura para representar a bola
type Bola struct {
    X, Y float64
    VX, VY float64
}

// Estrutura para representar a raquete
type Raquete struct {
    X, Y float64
    Altura float64
}

// Função principal
func main() {
    // Inicializando o gerador de números aleatórios
    rand.Seed(time.Now().UnixNano())

    // Criando a bola
    bola := Bola{
        X: LarguraTela / 2,
        Y: AlturaTela / 2,
        VX: rand.Float64() * 2 - 1,
        VY: rand.Float64() * 2 - 1,
    }

    // Criando as raquetes
    raqueteEsquerda := Raquete{
        X: 10,
        Y: AlturaTela / 2,
        Altura: 100,
    }
    raqueteDireita := Raquete{
        X: LarguraTela - 10,
        Y: AlturaTela / 2,
        Altura: 100,
    }

    // Loop principal do jogo
    for {
        // Atualizando a posição da bola
        bola.X += bola.VX
        bola.Y += bola.VY

        // Verificando se a bola saiu da tela
        if bola.X < 0 || bola.X > LarguraTela {
            // Reiniciando a bola
            bola.X = LarguraTela / 2
            bola.Y = AlturaTela / 2
            bola.VX = rand.Float64() * 2 - 1
            bola.VY = rand.Float64() * 2 - 1
        }
        if bola.Y < 0 || bola.Y > AlturaTela {
            // Invertendo a direção vertical da bola
            bola.VY = -bola.VY
        }

        // Verificando se a bola colidiu com uma raquete
        if bola.X < raqueteEsquerda.X+10 && bola.Y >= raqueteEsquerda.Y && bola.Y <= raqueteEsquerda.Y+raqueteEsquerda.Altura {
            // Invertendo a direção horizontal da bola
            bola.VX = -bola.VX
        }
        if bola.X > raqueteDireita.X-10 && bola.Y >= raqueteDireita.Y && bola.Y <= raqueteDireita.Y+raqueteDireita.Altura {
            // Invertendo a direção horizontal da bola
            bola.VX = -bola.VX
        }

        // Limpando a tela
        fmt.Print("\033[2J")

        // Desenhando a bola
        fmt.Printf("O\n")

        // Desenhando as raquetes
        fmt.Printf("|%s|\n", strings.Repeat(" ", int(raqueteEsquerda.Altura)))
        fmt.Printf("|%s|\n", strings.Repeat(" ", int(raqueteDireita.Altura)))

        // Aguardando um pouco antes de atualizar a tela
        time.Sleep(10 * time.Millisecond)
    }
}
```

