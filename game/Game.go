package game

import (
	"fmt"
	"github.com/DonMatano/flapy/config"
	"github.com/DonMatano/flapy/game/entities"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	bird                          *entities.Bird
	obstacles                     []*entities.Obstacle
	background                    *rl.Texture2D
	defaultObstacleWidth          float32
	defaultObstacleStartPosTop    rl.Vector2
	defaultObstacleStartPosBottom rl.Vector2
	obstacleMinHeight             float32
	obstacleMaxHeight             float32
	obstacleSpawnRate             float64
	lastObstacleSpawnTime         float64
	LastObstacleWasTop            bool
}

var gameInstance *Game

func GetGameInstance() *Game {
	if gameInstance == nil {

		birdPosition := rl.NewVector2(float32(config.WindowWidth)*0.1, float32(config.WindowHeight)*0.1)
		bird := entities.NewBird(birdPosition)
		bgTexture := rl.LoadTexture("resources/background.png")
		obstacle := createNewObstacle(
			rl.NewVector2(
				float32(config.WindowWidth)*0.9,
				float32(config.WindowHeight)*-0.05,
			),
			float32(config.WindowWidth)*0.05,
			float32(config.WindowHeight)*0.5,
			true,
		)
		defaultObstacleStartPositionTop := rl.NewVector2(
			float32(config.WindowWidth)*1.2,
			float32(config.WindowHeight)*-0.1,
		)
		defaultObstacleStartPositionBottom := rl.NewVector2(
			float32(config.WindowWidth)*1.2,
			float32(config.WindowHeight)*1,
		)
		return &Game{
			bird:                          bird,
			background:                    &bgTexture,
			obstacles:                     []*entities.Obstacle{obstacle},
			defaultObstacleWidth:          float32(config.WindowWidth) * 0.05,
			obstacleMinHeight:             float32(config.WindowHeight) * 0.3,
			obstacleMaxHeight:             float32(config.WindowHeight) * 0.8,
			obstacleSpawnRate:             5,
			lastObstacleSpawnTime:         rl.GetTime(),
			LastObstacleWasTop:            true,
			defaultObstacleStartPosTop:    defaultObstacleStartPositionTop,
			defaultObstacleStartPosBottom: defaultObstacleStartPositionBottom,
		}
	}
	return gameInstance
}

func createNewObstacle(position rl.Vector2, obstacleWidth, obstacleHeight float32, isTop bool) *entities.Obstacle {
	return entities.NewObstacle(position, obstacleWidth, obstacleHeight, isTop)
}

func (g *Game) shouldSpawnObstacle() bool {
	return (rl.GetTime() - g.lastObstacleSpawnTime) > g.obstacleSpawnRate
}

func (g *Game) spawnObstacle() {
	heightDiff := int32(g.obstacleMaxHeight - g.obstacleMinHeight)
	obstacleHeight := g.obstacleMinHeight + float32(rl.GetRandomValue(0, heightDiff))
	obstacle := &entities.Obstacle{}
	randomValue := rl.GetRandomValue(0, 1)
	isTop := randomValue == 1
	fmt.Println("Last obstacle was top: ", g.LastObstacleWasTop)
	if !isTop {
		obstacle = createNewObstacle(
			g.defaultObstacleStartPosBottom,
			g.defaultObstacleWidth,
			-obstacleHeight,
			false,
		)
		g.LastObstacleWasTop = false

	} else {
		obstacle = createNewObstacle(
			g.defaultObstacleStartPosTop,
			g.defaultObstacleWidth,
			obstacleHeight,
			true,
		)
		g.LastObstacleWasTop = true
	}
	g.obstacles = append(g.obstacles, obstacle)
}

func (g *Game) updateObstacles() {
	for _, obstacle := range g.obstacles {
		obstacle.Update()
	}
}

func (g *Game) drawObstacles() {
	for _, obstacle := range g.obstacles {
		obstacle.Draw()
	}
}

func (g *Game) Loop() {
	// Draw bg
	rl.DrawTexture(*g.background, 0, 0, rl.White)
	g.bird.Update()
	g.updateObstacles()
	g.drawObstacles()
	g.bird.Draw()
	// Check if we should spawn a new obstacle
	if g.shouldSpawnObstacle() {
		g.spawnObstacle()
		g.lastObstacleSpawnTime = rl.GetTime()
	}
	// clear obstacles that are out of screen
	for i, obstacle := range g.obstacles {
		if obstacle.IsOutOfScreen() {
			g.obstacles = append(g.obstacles[:i], g.obstacles[i+1:]...)
		}
	}

	fmt.Println("Obstacles length: ", len(g.obstacles))
	// Multiply by 0.95 and 0.05 to get the position of the FPS counter percentage-wise on the screen ie 95% from the left and 5% from the top
	rl.DrawFPS(int32(float32(config.WindowWidth)*0.95), int32(float32(config.WindowHeight)*0.05))
}
