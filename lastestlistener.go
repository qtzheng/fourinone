package fourinone

type EventListener interface {
}
type LastestListener interface {
	happenLastest(le LastestEvent) bool
}
