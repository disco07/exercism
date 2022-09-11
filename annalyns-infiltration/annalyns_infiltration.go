package annalyn

// CanFastAttack can be executed only when the knight is sleeping
func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	}
	return true
}

// CanSpy can be executed if at least one of the characters is awake
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	var knightIsAwakeVar int
	var archerIsAwakeVar int
	var prisonerIsAwakeVar int
	if knightIsAwake {
		knightIsAwakeVar = 1
	}
	if archerIsAwake {
		archerIsAwakeVar = 1
	}
	if prisonerIsAwake {
		prisonerIsAwakeVar = 1
	}
	switch knightIsAwakeVar + archerIsAwakeVar + prisonerIsAwakeVar {
	case 0:
		return false
	case 1, 2, 3:
		return true
	default:
		return false
	}
}

// CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if archerIsAwake || !prisonerIsAwake {
		return false
	}
	return true
}

// CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// or if Annalyn's pet dog is with her and the archer is sleeping
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	switch petDogIsPresent {
	case true:
		if !archerIsAwake {
			return true
		}
		return false
	case false:
		if !prisonerIsAwake {
			return false
		} else if prisonerIsAwake && !knightIsAwake && !archerIsAwake {
			return true
		}
	default:
		return false
	}
	return false
}
