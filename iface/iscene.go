package iface

type TileType int

const (
	Empty TileType = iota
	Wall
	Player
	Treasure
	Thunder
)
