package game

import "time"

type Config struct {
	BoardWidth      int
	BoardHeight     int
	StartLevel      int
	LinesPerLevel   int
	ScoreByLines    map[int]int
	BaseGravity     time.Duration
	GravityStep     time.Duration
	MinimumGravity  time.Duration
	SpawnX          int
	SpawnY          int
	WallKickOffsets []Point
}

func DefaultConfig() Config {
	return Config{
		BoardWidth:    10,
		BoardHeight:   20,
		StartLevel:    1,
		LinesPerLevel: 10,
		ScoreByLines: map[int]int{
			1: 100,
			2: 300,
			3: 500,
			4: 800,
		},
		BaseGravity:     800 * time.Millisecond,
		GravityStep:     60 * time.Millisecond,
		MinimumGravity:  120 * time.Millisecond,
		SpawnX:          3,
		SpawnY:          0,
		WallKickOffsets: []Point{{X: 0, Y: 0}, {X: -1, Y: 0}, {X: 1, Y: 0}, {X: -2, Y: 0}, {X: 2, Y: 0}},
	}
}

func (c Config) GravityForLevel(level int) time.Duration {
	if level < c.StartLevel {
		level = c.StartLevel
	}

	steps := level - c.StartLevel
	gravity := c.BaseGravity - (time.Duration(steps) * c.GravityStep)
	if gravity < c.MinimumGravity {
		return c.MinimumGravity
	}
	return gravity
}
