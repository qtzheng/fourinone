package fourinone

const (
	OP_READ AuthPolicy = 2<<iota - 1
	OP_READ_WRITE
	OP_ALL
)

type AuthPolicy uint

func AuthIncluded(targetAuth, curAuth AuthPolicy) bool {
	return (targetAuth & curAuth) == targetAuth
}
