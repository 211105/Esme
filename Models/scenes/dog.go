// dog.go
package scenes

import (
	
	"github.com/faiface/pixel"
	_ "image/jpeg"
	_ "image/png"
)

// Dog representa un perro en el juego
type Dog struct {
	X, Y, Vy float64
	Sprite   *pixel.Sprite
}

// setDogSprite establece el sprite del perro según el estado
func setDogSprite(dog *Dog, spriteMap map[string]*pixel.Sprite, state string) {
	switch state {
	case "Right1":
		dog.Sprite = spriteMap["dog2"]
	case "Right2":
		dog.Sprite = spriteMap["dog3"]
	case "Left1":
		dog.Sprite = spriteMap["dog5"]
	case "Left2":
		dog.Sprite = spriteMap["dog6"]
	case "JumpingRight":
		dog.Sprite = spriteMap["dog7"]
	case "JumpingLeft":
		dog.Sprite = spriteMap["dog8"]
	case "NeutralLeft":
		dog.Sprite = spriteMap["dog4"]
	default:
		dog.Sprite = spriteMap["dog1"]
	}
}

// Puedes agregar más funciones relacionadas con Dog aquí si es necesario
