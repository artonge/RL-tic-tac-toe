package main

import (
	"math/rand"
	"time"
)

// Associate to a board state a value
type state struct {
	boardState board
	value      float64
	count      int
}

// Describe an action of the agent
// aka where the agent put an x
type action struct {
	i, j int
}

// An agent that will learn to play tic-tac-toe
type agent struct {
	states  []*state // States list of all the games
	history []*state // States list of the current game
	sign    caseType // The sign use to play
	wins    int      // Number of wins
}

var r *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}

// The policy
func (a *agent) policy(s *state) action {
	// fmt.Println("	Running policy")
	var nextState *state
	var nextAction action
	posStates, posActions := a.possibleNextStates(s)
	for i, s := range posStates {
		if s.value > s.value && r.Intn(100) > 20 {
			nextState = s
			nextAction = posActions[i]
		}
	}
	if nextState == nil {
		i := r.Intn(len(posStates))
		nextAction = posActions[i]
	}
	return nextAction
}

// Return the possible states from a given states
func (a *agent) possibleNextStates(s *state) ([]*state, []action) {
	// fmt.Println("	Getting possible next states from\n", s.boardState)
	posStates := []*state{}
	posActions := []action{}
	var posBoard board
	// Look for empty cases on the board
	for i, l := range s.boardState {
		for j, c := range l {
			if c == e {
				// When a case is empty, add the associated state to posStates
				posBoard = s.boardState.copy()
				posBoard[i][j] = a.sign
				// fmt.Println("		", posBoard)
				posStates = append(posStates, a.getState(posBoard))
				posActions = append(posActions, action{i, j})
			}
		}
	}
	return posStates, posActions
}

// Called when we need the agent to play
func (a *agent) play(b board) {
	// fmt.Printf("%v plays\n", a.sign)
	s := a.getState(b)
	action := a.policy(s)
	a.history = append(a.history, s)
	b[action.i][action.j] = a.sign
}

// Called when the agent wins
func (a *agent) win() {
	a.wins++
	a.updateHistory(1)
}

// Called when the agent loses
func (a *agent) lose() {
	a.updateHistory(0)
}

func (a *agent) updateHistory(reward int) {
	// fmt.Printf("	Update %v's history\n", a.sign)
	for _, s := range a.history {
		s.value, s.count = incrementalMean(s.value, reward, s.count)
	}
	a.history = []*state{}
}

// Return a new mean value given old mean, new a value and the count of values involved
func incrementalMean(oldMean float64, value int, k int) (float64, int) {
	// fmt.Println("	Computing incremental mean")
	return (oldMean*float64(k) + float64(value)) / float64(k+1), k + 1
}

// Return a from the current board state
// If the agent have allready encounter the state, return it
// Else create a new one
func (a *agent) getState(b board) *state {
	// fmt.Println("	Getting state for board")
	var s *state
	// Search for an allready encounntered equivalent state
	// if len(a.states) > 0 {
	// 	fmt.Println("##########################")
	// 	fmt.Println(b)
	// 	fmt.Println("----")
	// 	fmt.Println(a.states[0].boardState)
	// }
	for _, state := range a.states {
		if state.boardState.isEqual(b) {
			// fmt.Println("Found state and value is ", state.value)
			// fmt.Println(b)
			s = state
			break
		}
	}
	// If state does not existe, create a new one and add it to states
	if s == nil {
		s = &state{
			value:      0,
			boardState: b.copy(),
		}
		a.states = append(a.states, s)
	}
	return s
}
