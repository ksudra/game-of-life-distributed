package main

import (
	"flag"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"time"

	"uk.ac.bris.cs/gameoflife/stubs"
	"uk.ac.bris.cs/gameoflife/util"
)

type GameOfLife struct{}

var aliveCount int
var turn int
var board [][]uint8
var shut bool
var pause bool
var quit bool
var finished bool

func (g *GameOfLife) GOL(request stubs.GameReq, response *stubs.GameRes) (err error) {
	tempWorld := make([][]uint8, len(request.World))
	for i := range request.World {
		tempWorld[i] = make([]uint8, len(request.World[i]))
		copy(tempWorld[i], request.World[i])
	}
	for i := 0; i < request.Turns; i++ {
		for pause {
		}
		aliveCount = len(calculateAliveCells(tempWorld))
		turn = i
		board = tempWorld
		tempWorld = calculateNextState(request.Width, request.Height, tempWorld)
		if shut {
			time.Sleep(200 * time.Millisecond)
			os.Exit(0)
		}
	}
	board = tempWorld
	response.World = board
	response.CompletedTurns = request.Turns
	response.Alive = calculateAliveCells(board)
	finished = true
	return
}

func calculateNextState(width, height int, world [][]uint8) [][]uint8 {
	tempWorld := make([][]uint8, len(world))
	for i := range world {
		tempWorld[i] = make([]uint8, len(world[i]))
		copy(tempWorld[i], world[i])
	}

	for y := range tempWorld {
		for x := range tempWorld {
			count := countNeighbours(width, height, x, y, world)

			if world[y][x] == 255 && (count < 2 || count > 3) {
				tempWorld[y][x] = 0
			} else if world[y][x] == 0 && count == 3 {
				tempWorld[y][x] = 255
			}
		}
	}

	return tempWorld
}

func countNeighbours(width, height, x, y int, world [][]uint8) int {
	neighbours := [8][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	count := 0

	for _, r := range neighbours {
		if world[(y+r[0]+height)%height][(x+r[1]+width)%width] == 255 {
			count++
		}
	}

	return count
}

func calculateAliveCells(world [][]uint8) []util.Cell {
	var cells []util.Cell
	for i := range world {
		for j := range world[i] {
			if world[i][j] == 255 {
				cells = append(cells, util.Cell{X: j, Y: i})
			}
		}
	}
	return cells
}

func (g *GameOfLife) GetNumAlive(request stubs.AliveReq, response *stubs.AliveRes) (err error) {
	response.Turn = turn
	response.Alive = aliveCount
	return
}

func (g *GameOfLife) StateChange(request stubs.ChangeStateReq, response *stubs.ChangeStateRes) (err error) {
	response.Turn = turn
	return
}

func (g *GameOfLife) GetBoard(request stubs.BoardReq, response *stubs.BoardRes) (err error) {
	response.Turn = turn
	response.Alive = calculateAliveCells(board)
	response.World = board
	return
}

func (g *GameOfLife) ShutDown(request stubs.CloseReq, response *stubs.CloseRes) (err error) {
	shut = true
	return
}

func (g *GameOfLife) PauseGame(request stubs.PauseReq, response *stubs.PauseRes) (err error) {
	if pause {
		pause = false
	} else {
		pause = true
	}
	response.Turn = turn
	return
}

func (g *GameOfLife) QuitGame(request stubs.QuitReq, response *stubs.QuitRes) (err error) {
	quit = true
	return
}
func (g *GameOfLife) CheckQuit(request stubs.CheckQuitReq, response *stubs.CheckQuitRes) (err error) {
	response.Quit = quit
	return
}

func (g *GameOfLife) CheckFinished(request stubs.FinishedReq, response *stubs.FinishedRes) (err error) {
	response.Finished = finished
	return
}

func main() {
	pAddr := flag.String("port", "8030", "Port to listen on")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	err := rpc.Register(&GameOfLife{})
	if err != nil {
		return
	}

	listener, _ := net.Listen("tcp", ":"+*pAddr)

	defer listener.Close()
	rpc.Accept(listener)
}
