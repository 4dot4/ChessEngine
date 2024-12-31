package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Piece int

const (
	Pawn Piece = iota
	Knight
	Rook
	Bisp
	Queen
	King
)

type Cell struct {
	Piece Piece
	Color rl.Color
}

func main() {
	rl.InitWindow(600, 600, "Maths chess")
	var invert = -1

	var board [8][8]Cell
	for y, _ := range board {

		for x, _ := range board[y] {
			switch invert {
			case -1:
				if x%2 == 0 {
					board[y][x].Color = rl.White
				} else {
					board[y][x].Color = rl.Black
				}
			case 1:

				if x%2 == 0 {
					board[y][x].Color = rl.Black
				} else {
					board[y][x].Color = rl.White
				}

			}

		}
		invert = -invert
	}
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		for y, _ := range board {
			for x, v := range board[y] {
				rl.DrawRectangleRec(rl.Rectangle{float32(x * 75), float32(y * 75), 75, 75}, v.Color)
			}
		}

		rl.EndDrawing()
	}
	rl.CloseWindow()
}
