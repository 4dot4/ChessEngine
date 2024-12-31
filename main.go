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
	SCALE float32   = 3
	Pawn  typePiece = iota
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

func resetBoardColors() {

}
func initBoard() [8][8]Cell {
	var invert = -1
	var board [8][8]Cell
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
	board[7][0].Piece = Piece{Rook, White}
	board[7][1].Piece = Piece{Knight, White}
	board[7][2].Piece = Piece{Bisp, White}
	board[7][3].Piece = Piece{Queen, White}
	board[7][4].Piece = Piece{King, White}
	board[7][5].Piece = Piece{Bisp, White}
	board[7][6].Piece = Piece{Knight, White}
	board[7][7].Piece = Piece{Rook, White}
	for i := 0; i < 8; i++ {
		board[6][i].Piece = Piece{Pawn, White}
	}
	return board
}
func main() {
	rl.InitWindow(600, 600, "Maths chess")
	spritesPositions := map[Piece]rl.Rectangle{
		{Pawn, White}:   {0, 0, 16, 32},
		{Knight, White}: {16, 0, 16, 32},
		{Rook, White}:   {16 * 2, 0, 16, 32},
		{Bisp, White}:   {16 * 3, 0, 16, 32},
		{Queen, White}:  {16 * 4, 0, 16, 32},
		{King, White}:   {16 * 5, 0, 16, 32},

		{Pawn, Black}:   {0, 0, 16, 32},
		{Knight, Black}: {16, 0, 16, 32},
		{Rook, Black}:   {16 * 2, 0, 16, 32},
		{Bisp, Black}:   {16 * 3, 0, 16, 32},
		{Queen, Black}:  {16 * 4, 0, 16, 32},
		{King, Black}:   {16 * 5, 0, 16, 32},
	}
	board := initBoard()
	blackSprites := rl.LoadTexture("sprites/BlackPieces.png")
	whiteSprites := rl.LoadTexture("sprites/WhitePieces.png")
	rl.SetTargetFPS(60)
	holding := struct {
		State bool
		Piece Piece
	}{
		State: false,
	}

	for !rl.WindowShouldClose() {

		for y, _ := range board {
			for x, v := range board[y] {
				if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
					if holding.State {
						if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{float32(x) * 75, float32(y) * 75, 75, 75}) {
							board[y][x].Piece = holding.Piece
							holding.State = false
						}
					}
					if rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{float32(x) * 75, float32(y) * 75, 75, 75}) {
						holding.Piece = v.Piece
						holding.State = true
					}
				}
			}
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for y, _ := range board {
			for x, v := range board[y] {
				rl.DrawRectangleRec(rl.Rectangle{float32(x) * 75, float32(y) * 75, 75, 75}, v.Color)
				switch v.Piece.ColorPiece {
				case White:
					rl.DrawTexturePro(whiteSprites, spritesPositions[v.Piece], rl.Rectangle{float32(x)*75 + 12, float32(y)*75 - 20, 16 * SCALE, 32 * SCALE}, rl.Vector2{0, 0}, 0, rl.RayWhite)
				case Black:
					rl.DrawTexturePro(blackSprites, spritesPositions[v.Piece], rl.Rectangle{float32(x)*75 + 12, float32(y)*75 - 20, 16 * SCALE, 32 * SCALE}, rl.Vector2{0, 0}, 0, rl.RayWhite)

				}
				// rl.DrawTexturePro(blackSprites, spritesPositions[v.Piece], rl.Rectangle{float32(x) * 75, float32(y) * 75, 16 * SCALE, 32 * SCALE}, rl.Vector2{0, 0}, 0, rl.RayWhite)
			}
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}
