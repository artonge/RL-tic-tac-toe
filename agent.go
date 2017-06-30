package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ---
// STATE DEFINITION
// ---
// Associate a board state to a value
type state struct {
	boardState board   // The board state
	value      float64 // The value associated to the state
	count      int     // The number of time the agent has seen this state
}

// Override String method
func (s state) String() string {
	return fmt.Sprintf("%v	{value: %v - count: %v}\n---", s.boardState, s.value, s.count)
}

// ---
// ACTION DEFINITION
// ---
// Describe an action of the agent
// aka where the agent put his sign
type action struct {
	i, j int
}

// ---
// AGENT DEFINITION
// ---
// An agent that will learn to play tic-tac-toe
type agent struct {
	history   map[string]*state // States encountered during of all the games
	gameMoves map[string]*state // States encountered during the current game
	sign      boardCase         // The sign use to play
}

// Called when we need the agent to play
func (a *agent) play(b board) {
	// Save the current state in gameMoves
	s := a.getState(b)
	a.gameMoves[s.boardState.serialize()] = s
	// Get the action from the policy
	act := a.policy(b)
	// Apply the action
	b[act.i][act.j] = a.sign
	// Save the new state in gameMoves
	s = a.getState(b)
	a.gameMoves[s.boardState.serialize()] = s
}

// Init random with time to have more randomness
var r *rand.Rand

func init() {
	s := rand.NewSource(time.Now().UnixNano())
	r = rand.New(s)
}

// The policy
// Given a state, it return an action
func (a *agent) policy(b board) action {
	maxVal := 0.0
	var nextAction action
	posStates, posActions := a.possibleNextStates(b)
	// Get the hightest valued state from posStates
	for i, s := range posStates {
		if s.value > maxVal {
			maxVal = s.value
			nextAction = posActions[i]
		}
	}
	// 10% of the time chose random actions
	if float64(r.Intn(100)) < 10 {
		nextAction = posActions[r.Intn(len(posActions))]
	}
	return nextAction
}

// Return the possible next states from a given states and the action to get there
func (a *agent) possibleNextStates(b board) ([]*state, []action) {
	posStates := []*state{}
	posActions := []action{}
	var posBoard board
	// Look for empty cases on the board
	for i, l := range b {
		for j, c := range l {
			if c == e {
				// When a case is empty, add the associated state to posStates
				posBoard = b.copy()
				posBoard[i][j] = a.sign
				posStates = append(posStates, a.getState(posBoard))
				posActions = append(posActions, action{i, j})
			}
		}
	}
	return posStates, posActions
}

// Called at the end of the game to feed the agent with the reward
func (a *agent) feed(reward int) {
	for _, s := range a.gameMoves {
		// Compute the incremental mean
		s.value = (s.value*float64(s.count) + float64(reward)) / float64(s.count+1)
		s.count++
	}
	a.gameMoves = make(map[string]*state)
}

// Return a state from the history or create a new one it's a board the agent has never seen
func (a *agent) getState(b board) *state {
	var s *state
	var ok bool
	// Check in history if agent has seen the board
	for i := 0; i < 4; i++ {
		s, ok = a.history[b.serialize()]
		if ok {
			return s
		}
		b = b.rotate()
	}
	if s == nil {
		// If the board has never been seen, create a new one and add it to history
		s = &state{
			value:      0,
			boardState: b.copy(),
		}
		a.history[s.boardState.serialize()] = s
	}
	return s
}
