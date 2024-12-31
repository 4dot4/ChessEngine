package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type typePiece int
type ColorPiece int
type Piece struct {
	typePiece  typePiece
	ColorPiece ColorPiece
}

const (
	Pawn typePiece = iota
	Knight
	Rook
	Bisp
	Queen
	King
	White ColorPiece = iota
	Black
)

type Cell struct {
	Piece Piece
	Color rl.Color
}

func main() {
	rl.InitWindow(533, 533, "Maths chess")
	spritesPositions := map[Piece]rl.Rectangle{
		Piece{King, White}:   rl.Rectangle{0, 0, 66.66, 64.5},
		Piece{Queen, White}:  rl.Rectangle{66.66, 0, 66.66, 64.5},
		Piece{Bisp, White}:   rl.Rectangle{66.66 * 2, 0, 66.66, 64.5},
		Piece{Knight, White}: rl.Rectangle{66.66 * 3, 0, 66.66, 64.5},
		Piece{Rook, White}:   rl.Rectangle{66.66 * 4, 0, 66.66, 64.5},
		Piece{Pawn, White}:   rl.Rectangle{66.66*5 - 10, 0, 66.66, 64.5},

		Piece{King, Black}:   rl.Rectangle{0, 64.5, 66.66, 64.5},
		Piece{Queen, Black}:  rl.Rectangle{66.66, 64.5, 66.66, 64.5},
		Piece{Bisp, Black}:   rl.Rectangle{66.66 * 2, 64.5, 66.66, 64.5},
		Piece{Knight, Black}: rl.Rectangle{66.66 * 3, 64.5, 66.66, 64.5},
		Piece{Rook, Black}:   rl.Rectangle{66.66 * 4, 64.5, 66.66, 64.5},
		Piece{Pawn, Black}:   rl.Rectangle{66.66 * 5, 64.5, 66.66, 64.5},
	}
	var invert = -1
	var board [8][8]Cell
	sprites := rl.LoadTexture("sprites/pieces.png")
	for y, _ := range board {

		for x, _ := range board[y] {
			switch invert {
			case -1:
				if x%2 == 0 {
					board[y][x].Color = rl.White
				} else {
					board[y][x].Color = rl.Gray
				}
			case 1:

				if x%2 == 0 {
					board[y][x].Color = rl.Gray
				} else {
					board[y][x].Color = rl.White
				}

			}
		}
		invert = -invert
	}
	board[0][0].Piece = Piece{Rook, Black}
	board[0][1].Piece = Piece{Knight, Black}
	board[0][2].Piece = Piece{Bisp, Black}
	board[0][3].Piece = Piece{Queen, Black}
	board[0][4].Piece = Piece{King, Black}
	board[0][5].Piece = Piece{Bisp, Black}
	board[0][6].Piece = Piece{Knight, Black}
	board[0][7].Piece = Piece{Rook, Black}
	for i := 0; i < 8; i++ {
		board[1][i].Piece = Piece{Pawn, Black}
	}
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		for y, _ := range board {
			for x, _ := range board[y] {
				if rl.IsMouseButtonPressed(rl.MouseLeftButton) {

					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{float32(x) * 66.66, float32(y) * 66.66, 66.66, 66.66}) {
						board[y][x].Color = rl.Red
					}
				}
			}
		}
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for y, _ := range board {
			for x, v := range board[y] {
				rl.DrawRectangleRec(rl.Rectangle{float32(x) * 66.66, float32(y) * 66.66, 66.66, 66.66}, v.Color)
				rl.DrawTextureRec(sprites, spritesPositions[v.Piece], rl.Vector2{float32(x) * 66.66, float32(y) * 64.5}, rl.White)
			}
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
