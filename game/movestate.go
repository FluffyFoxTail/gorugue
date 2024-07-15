package game

type MoveState int

const (
	BeforePlayerMove MoveState = iota
	PlayerMove
	EnemyMove
)

func GetNextState(state MoveState) MoveState {
	switch state {
	case BeforePlayerMove:
		return PlayerMove
	case PlayerMove:
		return EnemyMove
	case EnemyMove:
		return BeforePlayerMove
	default:
		return PlayerMove
	}
}
