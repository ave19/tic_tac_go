package tictacgo

import ()

func filterBytes(potential []byte, banned []byte) (accepted []byte) {
	for _, p := range potential {
		if !isBanned(p, banned) {
			accepted = append(accepted, p)
		}
	}
	return accepted
}

func uniqueBytes(list []byte) (newList []byte) {
	set := make(map[byte]struct{})
	for _, value := range list {
		set[value] = struct{}{}
	}
	for key := range set {
		newList = append(newList, key)
	}
	return
}

func isBanned(potential byte, banned []byte) bool {
	for _, b := range banned {
		if potential == b {
			return true
		}
	}
	return false
}
