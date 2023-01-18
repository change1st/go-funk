package v2

type JoinFnc[T comparable] func(lx, rx []T) []T

// Join combines two collections using the given join method.
func Join[T comparable](larr, rarr []T, fnc JoinFnc[T]) []T {
	return fnc(larr, rarr)
}

// InnerJoin finds and returns matching data from two collections.
func InnerJoin[T comparable](lx, rx []T) []T {
	result := make([]T, 0, len(lx)+len(rx))
	rhash := sliceSet(rx)
	lhash := make(map[T]struct{}, len(lx))

	for i := 0; i < len(lx); i++ {
		_, ok := rhash[lx[i]]
		_, alreadyExists := lhash[lx[i]]
		if ok && !alreadyExists {
			lhash[lx[i]] = struct{}{}
			result = append(result, lx[i])
		}
	}
	return result
}

// OuterJoin finds and returns dissimilar data from two collections.
func OuterJoin[T comparable](lx, rx []T) []T {
	ljoin := LeftJoin(lx, rx)
	rjoin := RightJoin(lx, rx)

	result := make([]T, 0, len(ljoin)+len(rjoin))
	for i := 0; i < len(ljoin); i++ {
		result = append(result, ljoin[i])
	}
	for i := 0; i < len(rjoin); i++ {
		result = append(result, rjoin[i])
	}

	return result
}

// LeftJoin finds and returns dissimilar data from the first collection (left).
func LeftJoin[T comparable](lx, rx []T) []T {
	result := make([]T, 0, len(lx))
	rhash := sliceSet(rx)

	for i := 0; i < len(lx); i++ {
		_, ok := rhash[lx[i]]
		if !ok {
			result = append(result, lx[i])
		}
	}
	return result
}

// RightJoin finds and returns dissimilar data from the second collection (right).
func RightJoin[T comparable](lx, rx []T) []T { return LeftJoin(rx, lx) }

func sliceSet[T comparable](arr []T) map[T]struct{} {
	hash := map[T]struct{}{}
	for i := 0; i < len(arr); i++ {
		hash[arr[i]] = struct{}{}
	}
	return hash
}
