package balances

type opts uint8

const DEFAULT = 0

const (
	VIRTUAL opts = 1 << iota
	PASSIVE
	SYNCABLE
)

func (o opts) isVirtual() bool {
	return (o & VIRTUAL) != 0
}

func (o opts) isPassive() bool {
	return (o & PASSIVE) != 0
}

func (o opts) isActive() bool {
	return (o & PASSIVE) == 0
}

func (o opts) isSyncable() bool {
	return (o & SYNCABLE) != 0
}
